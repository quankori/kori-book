# Blockchain Structure

- ### Cấu tạo một block

  - Hash, previous hash, data, nonce

  - previous hash + data + nonce + block id + hash alogirthm = hash gửi lên node

  - genesis block (block đầu tiên), previous sẽ là 00000000

  - The nonce là số lần transaction thực hiện thành công của user đó, miner sẽ dò nonce cho đến khi khớp với mã hash hệ thống đã tạo ra

  - Trong blockchain SHA256 hash data và là mã hoá một chiều.

  - Mã hash sẽ được biên dịch qua số thập phân (hexadecimal to decimal)

- ### Tính chất blockchain

  - The avalanche effect. khi thay đổi nhỏ 1 chỗ mã hash sẽ tạo lại 1 cái hoàn toàn khác với mã hash ban đầu

  - Chịu được va chạm (Must withstand collisions). Byzantine Fault Tolerance

    - Immutable ledger.

    - Những ưu điểm so với sự chứng nhận thông thường

    - Distributed P2P network

    - When add one block to network, it's copied onto all of the network (Consensus Protocal )

  - Can hack if u can control 51% node in network. But u must more machine, more piece to do it.

  - Nếu một node trong mạng lưới bị hack, tất cả node khác sẽ khôi phục lại node đó với giá trị gốc.

  - Khi 1 mining đưa 1 block vào chain thành công, các node còn lại trong network sẽ update theo.

  - Tuy nhiên, một trong những nhược điểm của thời gian khai thác một block ngắn là sẽ có nhiều thợ mỏ đào ra 1 block trong những khoàng thời gian rất gần nhau nhưng chỉ có một block được tìm ra sớm nhất được đưa vào nhánh chính và block chậm hơn sẽ bị bỏ được gọi là orphaned blocks.
