package main

import (

	"crypto/ecdsa"
	"fmt"
	"log"
	"os"
	"math/big"
	"encoding/json"
	"strconv"
	"context"
	"math/rand"
	"time"
	"flag"
	"io/ioutil"

	token "contracts/token"
	uniswap "contracts/uniswap"
	util "contracts/util"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	//"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
)

var ( // stop unused import warnings
	_ = json.Number(1)
	_ = os.O_RDONLY
)

var generic_nobody_address = common.HexToAddress("0x2adc25665018aa1fe0e6bc666dac8fc2697ff9ba")

type EnvGenerator struct {
	CurrentCoinbase 	common.Address 					`json:"currentCoinbase"`
	CurrentDifficulty 	*math.HexOrDecimal256			`json:"currentDifficulty"`
	GasLimit 			math.HexOrDecimal64				`json:"currentGasLimit"`
	CurrentNumber 		math.HexOrDecimal64				`json:"currentNumber"`
	CurrentTimestamp 	math.HexOrDecimal64				`json:"CurrentTimestamp"`
}

func safe_mkdir(dir string) {
	err := os.MkdirAll(dir, os.FileMode(0777))
	if (err != nil) {
		log.Fatal(err)
	}
}

func make_env(blockNumber uint64) {
	env := EnvGenerator {
		CurrentCoinbase: 	generic_nobody_address,
		CurrentDifficulty:	math.NewHexOrDecimal256(10000),
		GasLimit:			math.HexOrDecimal64(1000000000),
		CurrentNumber:		math.HexOrDecimal64(blockNumber),
		CurrentTimestamp:	math.HexOrDecimal64(blockNumber * 10),
	}

	str, err := json.MarshalIndent(env, "", "\t")
	if (err != nil) {
		log.Fatal(err)
	}

	safe_mkdir("env/")

	filename := "env/" + strconv.FormatUint(blockNumber, 10) + ".json"

	file, err := os.Create(filename)

	if (err != nil) {
		log.Fatal(err)
	}

	_, err = file.WriteString(string(str))
	if (err != nil) {
		log.Fatal(err)
	}
	file.Close()

}

type AccountContext struct {
	sk *ecdsa.PrivateKey
	pk *ecdsa.PublicKey
	nonce int64
	FromAddress common.Address
	signer types.EIP155Signer
	gasPrice *big.Int
	gasLimit uint64
}

type TokenContext struct {
	token *token.ERC20Mintable
	address common.Address
}

func (ctx *TokenContext) Approve(amount int64, sender_account *AccountContext, approved_agent common.Address) *types.Transaction {
	tx, err := ctx.token.Approve(account_to_tx_opts(sender_account), approved_agent, big.NewInt(amount))

	if (err != nil) {
		log.Fatal(err)
	}
	sender_account.increment_nonce()

	return tx
}

type TokenPairContext struct {
	pair 			*uniswap.UniswapV2Pair
	pair_itf 		*util.UniswapPairWrapper
	address 		common.Address // address of actual uniswap pair, not caller interface
	itf_address 	common.Address
}

func (ctx *TokenPairContext) ProvideLiquidity(amount int64, sender_account *AccountContext) *types.Transaction {
	tx, err := ctx.pair_itf.Mint(account_to_tx_opts(sender_account), big.NewInt(amount), big.NewInt(amount))

	sender_account.increment_nonce()

	if (err != nil) {
		log.Fatal(err)
	}

	return tx
}

//checking to see what failed uniswap swaps actually do
func (ctx *TokenPairContext) DoFailSwap(sender_account * AccountContext) *types.Transaction {
	tx, err := ctx.pair_itf.Swap(account_to_tx_opts(sender_account), big.NewInt(0), big.NewInt(0), big.NewInt(00000), big.NewInt(0000))
	if (err != nil) {
		log.Fatal(err)
	}
	sender_account.increment_nonce() //huh? Shouldn't get here, hopefully
	return tx
}

func (ctx *TokenPairContext) DoSwap0To1(amount int64, sender_account *AccountContext) *types.Transaction {
	// r0* r1 = (r0 + amount) * (r1 - withdrawable)
	// r1 - (r0 r1) / (r0+amount) = withdrawable

	reserves, err := ctx.pair.GetReserves(gen_call_opts())
	if (err != nil) {
		log.Fatal(err)
	}

	r0 := reserves.Reserve0
	r1 := reserves.Reserve1

//	fmt.Println("reserves 0 to 1", r0, r1, amount)

	prod := big.NewInt(0)

	prod.Mul(r0, r1)

	new_r0 := big.NewInt(0)
	new_r0.Add(r0, big.NewInt(charge_tax(amount)))

	withdrawable := big.NewInt(0).Set(r1)

	prod.Quo(prod, new_r0)

	withdrawable.Sub(withdrawable, prod) // withdrawable = r1 - (r0r1/(r0+amount))

//	fee := big.NewInt(0).Set(withdrawable)

//	fee.Quo(fee, big.NewInt(200)) // charge a .5% fee, instead of .3%, so as to not care about dumb rounding errors

//	withdrawable.Sub(withdrawable, fee)

	tx, err := ctx.pair_itf.Swap(account_to_tx_opts(sender_account),  big.NewInt(amount), big.NewInt(0), big.NewInt(0), withdrawable)

	sender_account.increment_nonce()
	if (err != nil) {
		log.Fatal(err)
	}
	return tx
}

func charge_tax (amount int64) int64 {
	return (amount * 995 ) / 1000
}

func (ctx *TokenPairContext) DoSwap1To0(amount int64, sender_account *AccountContext) *types.Transaction {
	// r0* r1 = (r1 + amount) * (r0 - withdrawable)
	// r0 - (r0 r1) / (r1+amount) = withdrawable

	reserves, err := ctx.pair.GetReserves(gen_call_opts())
	if (err != nil) {
		log.Fatal(err)
	}

	r0 := reserves.Reserve0
	r1 := reserves.Reserve1



	prod := big.NewInt(0)

	prod.Mul(r0, r1)

	new_r1 := big.NewInt(0)
	new_r1.Add(r1, big.NewInt(charge_tax(amount)))

	withdrawable := big.NewInt(0).Set(r0)

	prod.Quo(prod, new_r1)

	withdrawable.Sub(withdrawable, prod) // withdrawable = r0 - (r0r1/(r1+amount))

	//fee := big.NewInt(0).Set(withdrawable)

	//fee.Quo(fee, big.NewInt(200)) // charge a .5% fee, instead of .3%, so as to not care about dumb rounding errors

	//withdrawable.Sub(withdrawable, fee)

	//fmt.Println("reserves 1 to 0", r0, r1, amount, withdrawable, new_r1)

	tx, err := ctx.pair_itf.Swap(account_to_tx_opts(sender_account),  big.NewInt(0), big.NewInt(amount), withdrawable, big.NewInt(0))

	sender_account.increment_nonce()
	if (err != nil) {
		log.Fatal(err)
	}
	return tx
}

func gen_random_swap(token_pairs *[]TokenPairContext, accounts *[]*AccountContext) types.Transaction {
	sender_idx := rand.Int() % len(*accounts)
	sender_account := (*accounts)[sender_idx]

	pair_idx := rand.Int() % len(*token_pairs)
	pair := (*token_pairs)[pair_idx]

	direction := rand.Int() % 2
	if (direction == 0) {
		return *pair.DoSwap0To1(int64(1000), sender_account)
	} else {
		return *pair.DoSwap1To0(int64(1000), sender_account)
	}
}

func gen_swap_txs(num_txs uint64, token_pairs *[]TokenPairContext, accounts *[]*AccountContext) []types.Transaction {
	txs := []types.Transaction{}

	for i := uint64(0); i < num_txs; i++ {
		fmt.Println(i)
		txs = append(txs, gen_random_swap(token_pairs, accounts))
	}
	return txs
}

func gen_call_opts() *bind.CallOpts {
	opts := bind.CallOpts {
		Pending : true,
		Context: context.Background(),
	}
	return &opts;
}

func authorize_uniswap_itfs(token0 *TokenContext, token1 *TokenContext,  pair_if *TokenPairContext, accounts *[]*AccountContext, backend bind.ContractBackend) []types.Transaction {
	txs := make([]types.Transaction, 2*len(*accounts), 2*len(*accounts))

	//kinda silly but oh well
	for i, account := range *accounts {
		//fmt.Println(i)
		txs[2*i] = *token0.Approve(1000000000, account, pair_if.itf_address)
		//txs = append(txs, *tx)
		txs[2*i + 1] = *token1.Approve(1000000000, account, pair_if.itf_address)
		//txs = append(txs, *tx)
	}
	return txs
}

func deploy_uniswap_contracts(tokens []TokenContext, accounts *[]*AccountContext, backend bind.ContractBackend) ([]TokenPairContext, []types.Transaction) {

	token_pairs := []TokenPairContext{}
	txs := []types.Transaction{}

	creator_account := (*accounts)[0]

	_, factory_create_tx, factory, err := uniswap.DeployUniswapV2Factory(account_to_tx_opts(creator_account), backend, generic_nobody_address)
	if (err != nil) {
		log.Fatal(err)
	}
	creator_account.increment_nonce()

	txs = append(txs, *factory_create_tx)

	for i := 0; i < len(tokens); i++ {
		for j := i+1; j < len(tokens); j++ {
			fmt.Println("pairs for token", i, j)

			create_tx, err := factory.CreatePair(account_to_tx_opts(creator_account), tokens[i].address, tokens[j].address)
			if (err != nil) {
				log.Fatal(err)
			}
			txs = append(txs, *create_tx)
			creator_account.increment_nonce()


			pair_addr, err := factory.GetPair(gen_call_opts(), tokens[i].address, tokens[j].address)
			if (err != nil) {
				log.Fatal(err)
			}

			pair, err := uniswap.NewUniswapV2Pair(pair_addr, backend)
			if (err != nil) {
				log.Fatal(err)
			}

			itf_addr, create_itf_tx, pair_itf, err := util.DeployUniswapPairWrapper(account_to_tx_opts(creator_account), backend, pair_addr, tokens[i].address, tokens[j].address)
			if (err != nil) {
				log.Fatal(err)
			}
			creator_account.increment_nonce()
			txs = append(txs, *create_itf_tx)


			pair_ctx := TokenPairContext {
				pair : pair,
				pair_itf : pair_itf,
				address : pair_addr,
				itf_address: itf_addr,
			}

			token_pairs = append(token_pairs, pair_ctx)

			auth_txs := authorize_uniswap_itfs(&tokens[i], &tokens[j], &pair_ctx, accounts, backend)
			txs = append(txs, auth_txs...)
		}
	}

	fmt.Println("deployed uniswap")

	for _, tok := range tokens {
		fmt.Println(tok.token.BalanceOf(gen_call_opts(), creator_account.FromAddress))
		fmt.Println(tok.token.Allowance(gen_call_opts(), creator_account.FromAddress, token_pairs[3].itf_address))
	}

	liquidity_txs := provide_all_liquidity(&token_pairs, creator_account)
	txs = append(txs, liquidity_txs...)
	return token_pairs, txs
}

func provide_all_liquidity(pairs *[]TokenPairContext, sender *AccountContext) []types.Transaction {

	txs := []types.Transaction{}

	for i, pair := range *pairs {
		fmt.Println("providing liquidity to pair ", i)
		tx := pair.ProvideLiquidity(100000, sender)
		txs = append(txs, *tx)
	}
	return txs

}

func make_new_account() *AccountContext {
	sk, err := crypto.GenerateKey()

	if (err != nil) {
		log.Fatal(err)
	}

	publicKey := sk.Public()

	pkECDSA := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*pkECDSA)

	return &AccountContext {
		sk:				sk,
		pk: 			pkECDSA,
		nonce: 			0,
		FromAddress: 	fromAddress,
		signer: 		types.NewEIP155Signer(big.NewInt(1337)),
		gasPrice: 		big.NewInt(1),
		gasLimit: 		10000000}
}

func (ac *AccountContext) increment_nonce() {
	ac.nonce ++
}

func account_to_tx_opts(account *AccountContext) *bind.TransactOpts {
	opts, err :=  bind.NewKeyedTransactorWithChainID(account.sk, big.NewInt(1337))
	if (err != nil) {
		log.Fatal(err)
	}
	opts.Nonce = big.NewInt(account.nonce)
	return opts
}

func make_accounts(num_accounts uint64) []*AccountContext {
	accounts_out := []*AccountContext{}
	for i:= uint64(0); i < num_accounts; i++ {
		//if (int64(i) % 100 == 0) {
		//	fmt.Println("making account ", i)
		//}
		accounts_out = append(accounts_out, make_new_account())
	}
	return accounts_out
}

func load_genesis(genesis_block_number uint64) core.GenesisAlloc {
	genesis_alloc := make(core.GenesisAlloc)

	data, err := ioutil.ReadFile("alloc/" + strconv.FormatUint(genesis_block_number, 10) + ".json")
	if (err != nil) {
		log.Fatal(err)
	}
	
	json.Unmarshal(data, &genesis_alloc)
	return genesis_alloc
}


func make_genesis(accounts []*AccountContext, db ethdb.Database, genesis_block_number uint64) (*backends.SimulatedBackend) {
	genesis_alloc := make(core.GenesisAlloc)

	for _, account := range accounts {
		genesis_alloc[account.FromAddress] = core.GenesisAccount{
			Balance: big.NewInt(10000000000),
			Nonce: uint64(account.nonce)}
	}

	stringify, _ := json.MarshalIndent(genesis_alloc, "", "\t")
	str := string(stringify)

	safe_mkdir("alloc/")

	file, err := os.Create("alloc/" + strconv.FormatUint(genesis_block_number, 10) + ".json")

	if (err != nil) {
		log.Fatal(err)
	}

	_, err = file.WriteString(str)
	if (err != nil) {
		log.Fatal(err)
	}

	file.Close()

	sim := backends.NewSimulatedBackendWithDatabase(db, genesis_alloc, 1000000000) // 2nd arg is gas limit

	return sim
}

func deploy_tokens(backend bind.ContractBackend, accounts []*AccountContext, num_tokens uint64) ([]TokenContext, []types.Transaction) {

	tokens := []TokenContext{}
	txs := []types.Transaction{}

	for i := uint64(0); i < num_tokens; i++ {
		name := "dumb_token_" + strconv.FormatUint(i, 10)
		symbol := "DT" + strconv.FormatUint(i, 10)
		fmt.Println("making token ", name)
		token_addr, creation_tx, tok, err := token.DeployERC20Mintable(account_to_tx_opts(accounts[0]), backend, name, symbol)
		accounts[0].increment_nonce()

		txs = append(txs, *creation_tx)

		tok_ctx := TokenContext {
			token : tok,
			address: token_addr,
		}

		tokens = append(tokens, tok_ctx)

		if (err != nil) {
			log.Fatal(err)
		}
		
		mint_txs := make([]types.Transaction, len(accounts), len(accounts))

		for j := 0; j < len(accounts); j++ {

			mint_tx, err := tok.Mint(account_to_tx_opts(accounts[j]), accounts[j].FromAddress, big.NewInt(1000000000000))
			if (err != nil) {
				log.Fatal(err)
			}
			mint_txs[j] = *mint_tx

//			txs = append(txs, *mint_tx)
			accounts[j].increment_nonce()
		}
		txs = append(txs, mint_txs...)
	}
	return tokens, txs
}

func dumpState(backend *backends.SimulatedBackend, blockNumber int) (string) {
	block, err := backend.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if (err != nil) {
		log.Fatal(err)
	}
	data, err := json.MarshalIndent(block, "", "\t")
	if (err != nil) {
		log.Fatal(err)
	}

	return string(data)
}

func dump_txs(txs []types.Transaction, blockNumber uint64) {
	stringify, _ := json.MarshalIndent(txs, "", "\t")
	str := string(stringify)

	safe_mkdir("txs/")

	file, err := os.Create("txs/" + strconv.FormatUint(blockNumber, 10) + ".json")

	if (err != nil) {
		log.Fatal(err)
	}

	_, err = file.WriteString(str)
	if (err != nil) {
		log.Fatal(err)
	}

	file.Close()
}


func gen_swap_block(token_pairs *[]TokenPairContext, accounts *[]*AccountContext, num_txs uint64, block_number *uint64, simulator *backends.SimulatedBackend) {

	fmt.Println("making block", *block_number)
	txs := gen_swap_txs(num_txs, token_pairs, accounts)

	dump_txs(txs, *block_number)
	make_env(*block_number)
	simulator.Commit()
	dump_block(*block_number, simulator)
	*block_number++
}

func clear_envs() {
	err := os.RemoveAll("./txs")
	if (err != nil) {
		log.Fatal(err)
	}
	err = os.RemoveAll("./env")
	if (err != nil) {
		log.Fatal(err)
	}
	err = os.RemoveAll("./alloc")
	if (err != nil) {
		log.Fatal(err)
	}
	err = os.RemoveAll("./block")
	if (err != nil) {
		log.Fatal(err)
	}
}

type BlockWrapper struct {
	Header 		types.Header
	Txs 		[]*types.Transaction
}

func serialize_block(block *types.Block) BlockWrapper {
	body := block.Body()
	//transactions := []types.Transaction{}
	return BlockWrapper {
		Header : *block.Header(),
		Txs : body.Transactions,
	}
}

func deserialize_block(wrapper BlockWrapper) *types.Block {
	block := types.NewBlockWithHeader(&wrapper.Header)
	uncles := []*types.Header{}
	//fmt.Println("deserialized:", wrapper.Txs)
	block = block.WithBody(wrapper.Txs, uncles)
	//fmt.Println(block.Transactions())
	return block
}

func load_block(block_number uint64) *types.Block {
	data, err := ioutil.ReadFile("block/" + strconv.FormatUint(block_number, 10) + ".json")
	if (err != nil) {
		log.Fatal(err)
	}

	wrapper := BlockWrapper{}
	json.Unmarshal(data, &wrapper)
	return deserialize_block(wrapper)
}

func dump_block(block_number uint64, simulator *backends.SimulatedBackend) {
	block, err := simulator.BlockByNumber(context.Background(), big.NewInt(int64(block_number)))
	if (err != nil) {
		log.Fatal(err)
	}

	safe_mkdir("block/")

	serializable_block := serialize_block(block)

	stringify, _ := json.MarshalIndent(serializable_block, "", "\t")
	str := string(stringify)

	file, err := os.Create("block/" + strconv.FormatUint(block_number, 10) + ".json")
	if (err != nil) {
		log.Fatal(err)
	}

	_, err = file.WriteString(str)
	if (err != nil) {
		log.Fatal(err)
	}

	file.Close()
}

// inclusive - start with alloc starting_block, run txs starting_block
// until we run txs ending_block, ending in alloc ending_block+1
func run_txs(ending_block uint64) { 
	db := rawdb.NewMemoryDatabase()

	alloc := load_genesis(0)

	genesis := core.Genesis {Config: params.AllEthashProtocolChanges, GasLimit: uint64(100000000), Alloc : alloc}
	genesis.MustCommit(db)

	bc, err := core.NewBlockChain(db, nil, params.AllEthashProtocolChanges, ethash.NewFaker(), vm.Config{}, nil, nil)
	if (err != nil) {
		log.Fatal(err)
	}

	genesis_block := load_block(0)

	bc.ResetWithGenesisBlock(genesis_block)

	//backend := backends.NewSimulatedBackendWithDatabase(db, alloc, uint64(10000000000))

	state_processor := bc.Processor()//core.NewStateProcessor(backend.Blockchain().Config(), backend.Blockchain(), backend.Blockchain().Engine())

	vm_cfg := bc.GetVMConfig()

	for i := uint64(1); i <= ending_block; i++ {

		fmt.Println("processing block", i)
		state_db , err := bc.State()
		if (err != nil) {
			log.Fatal(err)
		}

		block := load_block(i)
//		fmt.Println(block.Transactions())


		start := time.Now()
		receipts, logs, gas, err := state_processor.Process(block, state_db, *vm_cfg)
		if (err != nil) {
			log.Fatal(err)
		}
		duration := time.Since(start)

		status, err := bc.WriteBlockWithState(block, receipts, logs, state_db, false)

		if (err != nil) {
			log.Fatal(err)
		}

		if (status != core.CanonStatTy) {
			log.Fatal("didn't add block to canonical chain?")
		}

		//fmt.Println(status)
		//fmt.Println(bc.CurrentBlock().NumberU64())

		/*prev_td := bc.GetTdByHash(block.ParentHash())

		td := new(big.Int).Set(block.Difficulty())

		if (prev_td != nil) {
			td = td.Add(td, bc.GetTdByHash(block.ParentHash()))
		}

		rawdb.WriteTd(db, block.Hash(), block.NumberU64(), td)
		rawdb.WriteBlock(db, block)
		state_db.Commit(false)
		*/

		fmt.Println("duration :", duration.String());
		fmt.Println("gas used: ", gas);
		fmt.Println("TPS: ", (float64(len(block.Transactions())) / float64(duration.Milliseconds())) * 1000.0)
		fmt.Println(duration)
		fmt.Println(gas)
		//fmt.Println(receipts)

		//fmt.Println(bc.CurrentHeader().Number)
		//blk := bc.CurrentBlock()

		//fmt.Println(blk.Transactions())

	}
}

func gen_txs(num_blocks uint64, num_txs uint64, num_tokens uint64, num_accounts uint64) {

	clear_envs()

	accounts := make_accounts(num_accounts)

	db := rawdb.NewMemoryDatabase()

	current_block := uint64(0)

	simulator := make_genesis(accounts, db, current_block)

	tokens, txs := deploy_tokens(simulator, accounts, num_tokens)
	fmt.Println("deployed tokens")
	token_pairs, txs_add := deploy_uniswap_contracts(tokens, &accounts, simulator)
	txs = append(txs, txs_add...)

	dump_txs(txs, current_block)
	make_env(current_block)
	dump_block(current_block, simulator)
	simulator.Commit()
	fmt.Println("finished setup")
	current_block++


	for i := uint64(0); i < num_blocks; i++ {
		gen_swap_block(&token_pairs, &accounts, num_txs, &current_block, simulator)
	}
}

func main () {

	gen_cmd := flag.NewFlagSet("gen", flag.ExitOnError)
	gen_num_blocks := gen_cmd.Int("num_blocks", 0, "num_blocks")
	gen_txs_per_block := gen_cmd.Int("num_txs", 100, "num_txs")
	gen_num_tokens := gen_cmd.Int("num_tokens", 10, "num_tokens")
	gen_num_accounts := gen_cmd.Int("num_accounts", 10, "num_accounts")

	run_cmd := flag.NewFlagSet("run", flag.ExitOnError)
	//run_start_block := run_cmd.Int("start", 1, "start")
	run_end_block := run_cmd.Int("end", 0, "end")

	if len(os.Args) < 2 {
		log.Fatal("expected run or gen subcommands")
	}

	switch os.Args[1] {
	case "run":
		run_cmd.Parse(os.Args[2:])
		run_txs(uint64(*run_end_block))
	case "gen":
		gen_cmd.Parse(os.Args[2:])
		gen_txs(uint64(*gen_num_blocks), uint64(*gen_txs_per_block), uint64(*gen_num_tokens), uint64(*gen_num_accounts))
	default:
		log.Fatal("invalid subcommand")
	}
}



	//vm.WriteLogs(os.Stdout, state_db.Logs())

	//fmt.Println("what the fuck")


/*

	secretKey, err := crypto.GenerateKey()

	if (err != nil) {
		log.Fatal(err)
	}

	secretKeyBytes := crypto.FromECDSA(secretKey)

	fmt.Println(hexutil.Encode(secretKeyBytes))

	publicKey := secretKey.Public()

	pkECDSA := publicKey.(*ecdsa.PublicKey)

	pkBytes := crypto.FromECDSAPub(pkECDSA)
	fmt.Println(hexutil.Encode(pkBytes))


	fromAddress := crypto.PubkeyToAddress(*pkECDSA)
	fmt.Println(fromAddress)

	nonce := uint64(0)

	value := big.NewInt(100)
	gasLimit := uint64(1000)

	gasPrice := big.NewInt(100)

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")


	var data []byte

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainId := big.NewInt(1337)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), secretKey)


	tx_out, err := json.MarshalIndent(&signedTx, "", "\t")

	os.Stdout.Write(tx_out)*/



//}