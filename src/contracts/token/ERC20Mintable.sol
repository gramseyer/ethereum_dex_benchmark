pragma solidity ^0.8.0;

import "../openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";


/*
 * Simplistic ERC20 token that any contract can mint to anyone else
*/
contract ERC20Mintable is ERC20 {
	
	constructor(string memory name, string memory symbol) ERC20(name, symbol) {}

	function mint(address to, uint256 amount) public virtual {
		_mint(to, amount);
	}
}