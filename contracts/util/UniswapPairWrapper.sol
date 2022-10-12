pragma solidity ^0.8.0;

import '../uniswap/interfaces/IUniswapV2Pair.sol';
import '../openzeppelin-contracts/contracts/token/ERC20/IERC20.sol';

contract UniswapPairWrapper {
	IUniswapV2Pair private pair_interface;

	IERC20 private token0;
	IERC20 private token1;


	constructor(address pair_addr, address tok0_addr, address tok1_addr) public {
		pair_interface = IUniswapV2Pair(pair_addr);
		token0 = IERC20(tok0_addr);
		token1 = IERC20(tok1_addr);
	}

	function mint(uint amount0, uint amount1) public {
		send_to_uniswap(amount0, amount1);
		pair_interface.mint(msg.sender);
	}

	function swap(uint amount0in, uint amount1in, uint amount0out, uint amount1out) public {
		send_to_uniswap(amount0in, amount1in);
		pair_interface.swap(amount0out, amount1out, msg.sender, "");
	}

	function send_to_uniswap(uint amount0, uint amount1) private {
		require(token0.transferFrom(msg.sender, address(pair_interface), amount0), 'TRANSFER FAILED');
		require(token1.transferFrom(msg.sender, address(pair_interface), amount1), 'TRANSFER_FAILED');
	}
}