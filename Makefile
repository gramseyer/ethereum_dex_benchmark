

CONTRACTS = \
	src/contracts/token/ERC20Mintable.sol \
	src/contracts/uniswap/UniswapV2Factory.sol \
	src/contracts/util/UniswapPairWrapper.sol

CONTRACTS_GO=${CONTRACTS:.sol=.go}

%.go : %.sol
	abigen \
		--sol=$< \
		--pkg=contracts \
		--out=$@

uniswap_sim : $(CONTRACTS_GO)
	go build src/uniswap_sim.go

all: uniswap_sim

clean:
	rm $(CONTRACTS_GO)


