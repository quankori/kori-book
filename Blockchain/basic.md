# Blockchain cơ bản

- Cấu tạo một block

Hash, previous hash, data, nonce (previous hash + data + nonce + block id + hash alogirthm = hash gửi lên node)

Khái niệm genesis block (block đầu tiên)
Lúc này previous sẽ là 00000000
SHA256 hash data trong blockchain
One way

- Deterministic (xác dịnh)
- Fast computation
- The avalanche effect. khi thay đổi nhỏ 1 chỗ mã hash sẽ tạo lại 1 cái hoàn toàn khác với mã hash ban đầu
- Must withstand collisions (chịu được va chạm Ex: Byzantine Fault Tolerance)

  - Immutable ledger
  - Những ưu điểm so với sự chứng nhận thông thường
  - Distributed P2P network
  - When add one block to network, it's copied onto all of the network (Consensus Protocal )

- Can hack if u can control 51% node in network. But u must more machine, more piece to do it
- If one node in network be hacked, all node different will restore that node. with orignal value (Byzantine Fault Tolerance)
  Mining (phần này khá phức tạp về cách mining giải một bài toán)
  The nonce là số lần transaction thực hiện thành công của user đó, miner sẽ dò nonce cho đến khi khớp với mã hash hệ thống đã tạo ra (SHA256 phần the avalance effect)
- Mã hash sẽ được biên dịch qua số thập phân (hexadecimal to decimal)
- Ở một số chain cơ chế sẽ là dò số nonce cho đến khi đạt được vào vùng target
- Vùng target được miner đặt ra là vùng khi số nonce dò được rơi vào sẽ được giải ( thường là 000000 ở đầu mã hash được giải qua hexadecimal to decima)
  Byzantine Fault Tolerance
  Bài toán các vị tướng Byzantine
- Nói về sự đồng thuận theo số đông và cơ chế hạn chế fake data.
Consensus Protocal (chưa đụng sâu vào thuật toán)
- A Consensus Protocol is a process used to achieve agreement between participants of a distributed network
- Ứng dụng mining và Byzantine Fault Tolerance để xoá bỏ rủi ro hack data
- Khi 1 mining đưa 1 block vào chain thành công, các node còn lại trong network sẽ update theo
- Tuy nhiên, một trong những nhược điểm của thời gian khai thác một block ngắn là sẽ có nhiều thợ mỏ đào ra 1 block trong những khoàng thời gian rất gần nhau nhưng chỉ có một block được tìm ra sớm nhất được đưa vào nhánh chính và block chậm hơn sẽ bị bỏ được gọi là orphaned blocks.