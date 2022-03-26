# Blockchain Quiz

## Tài khoản trong Ethereum

- Có 2 loại tài khoản trong Ethereum là smart contract và tài khoản thuộc sở hữu ( Externally-Owned Accounts )

## Sự khác nhau giữa transfer(), send() và call.value()

- Nếu transfer() gặp vấn đề, nó sẽ đưa ra ngoại lệ, nó sẽ khiến cho giao dịch bị hủy bỏ. Điều này thường sẽ chỉ xảy ra nếu quá trình chuyển hết gas. Hủy bỏ giao dịch trong trường hợp này là tốt và an toàn vì có lẽ bạn cũng không muốn giao dịch hoàn thành với giả định sai là nó đã thực hiện chuyển ether.

- send() sẽ trả về true hoặc false tùy thuộc vào việc nó chuyển ether thành công hay thất bại, nhưng nó sẽ không bao giờ hủy bỏ. Nếu smart contract không kiểm tra gía trị trả về của send() hoặc không xử lý được chính xác lỗi thì smart contract có thể rơi vào trạng thái không nhất quán và không thể khôi phục. Do đó, tôi khuyên bạn hãy sử dụng transfer() thay vì send() để chuyển tiền ra khỏi smart contract

- address.call.value(amount)( ).gas() đây là một low-level function và nó sẽ trả lại false nếu xảy ra lỗi . Sự khác biệt của nó so với hai chức năng trên là bạn phải set gas thông qua .gas(gasLimit) nếu không nó sẽ gây tốn gas trong những hợp đồng phải thực hiện một logic phức tạp , đòi hỏi nhiều gas

## Làm sao bảo vệ data quan trọng trong solidity

- Không nên xài private data, chỉ có thể sử dụng hashes

## Khác nhau giữa uint8 và uint16

- 2 ^ 8 - 1 hoặc 2 ^ 16 - 1

## Sản phẩm sau khi solidity compile

- abi, bytecode

## ABI của smart contract

- interface của contract

## Transaction nếu ko đủ gas

- transaction sẽ fail

## Làm sao quản lý time trong solidity

- Sử dụng block timestamp

## Solidity có mấy bộ nhớ lưu trữ

- Storage, memory, stack, calldata

## Làm sao để cancel 1 transaction đang pending

- Gửi 1 transaction khác cùng nonce và giá gasPrice cao hơn

## Khái niệm gas

- Gas Fee: tiền giao dịch khi người dùng phải trả để thực hiện giao dịch

- Gas Limit: Gas limit là lượng gas tối đa mà bạn sẽ chi trả cho một giao dịch, nếu bạn đặt gas limit thấp hơn số gas cần thiết để chạy giao dịch thì giao dịch này sẽ thất bại, bạn sẽ mất toàn bộ số gas fee sử dụng trong giao dịch này. Tuy nhiên thường thì các giao dịch không sử dụng hết toàn bộ số gas limit, bạn có thể giảm xuống nếu cần thiết để thực hiện giao dịch.

- Gas Price: Gas Price là số tiền tính bằng đồng coin gốc của blockchain đó mà người dùng sẵn sàng chi cho mỗi đơn vị Gas. Gas Price sẽ ảnh hưởng đến tốc độ xác nhận giao dịch của miner/validator trong network và đưa chúng vào block mới.

## Phân biệt layer 1 và layer 2

- Layer 1 Blockchain dùng để chỉ chuỗi chính của blockchain được hiểu đơn giản là kiến trúc blockchain ban đầu.

- Layer 2 là các giao thức hay nền tảng hỗ trợ blockchain ban đầu, với mục đích giúp blockchain đó giải quyết các vấn đề như khả năng mở rộng, tốc độ xử lý, phí,…
