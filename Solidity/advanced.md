# Solidity Advanced

- ### Tương tác qua một smart contract bất kỳ

```solidity
(bool status, bytes memory ret) = address_value.call{
    value: amount_value
}(data_value_bytes);
```

- ### Tương tác qua một smart contract được định nghĩa sẵn

```solidity
IERC20 token = IERC20(currency);
token.transferFrom(msg.sender, _admin, amount);
```
