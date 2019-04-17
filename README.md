# API-Create-Account-EOS
API Create Account EOS

Tạo account thông qua API POST api/v1/account/create, body {"AccountName":"example account name"}  
Go server run file script tạo acccount: createAccount.sh (hãy đảm bảo rằng keosd và nodeos đang chạy)  
File txt chứa cặp public key và private key (1 key pair tạo 1 account): keyPair.txt  
File txt chứa thông tin về account name, wallet name, wallet password: information_need_for_create_account.txt (Account name lấy từ POST request, wallet name và wallet password nếu POST request không truyền vào thì lấy mặc định)
Tất cả các file trên đặt ở thư mục ~/eos/build/programs/cleos.
