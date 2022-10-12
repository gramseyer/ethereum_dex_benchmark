

CONTRACTS = \
	contracts/token/ERC20Mintable.sol \
	contracts/uniswap/UniswapV2Factory.sol \
	contracts/univ2pair/UniswapV2Pair.sol \
	contracts/util/UniswapPairWrapper.sol

CONTRACTS_GO=${CONTRACTS:.sol=.go}

%.go : %.sol
	solc --bin -o $(basename $<).bin --overwrite $< && \
	solc --abi -o $(basename $<).abi --overwrite $< && \
	abigen \
		--bin $(basename $<).bin/$(basename $(notdir $<)).bin \
		--abi $(basename $<).abi/$(basename $(notdir $<)).abi \
		--pkg=$(basename $(notdir $<)) \
		--out=$@

#$(notdir $(patsubst %/,%,$(basename $(dir $<))))


uniswap_sim : $(CONTRACTS_GO) $(CONTRACTS) uniswap_sim.go
	go build uniswap_sim.go

all: uniswap_sim

clean:
	rm $(CONTRACTS_GO)


