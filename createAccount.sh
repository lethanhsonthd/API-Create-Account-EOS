#!/bin/bash
#!/usr/bin/expect
EOS_CLEOS_PATH="/home/son/eos/build/programs/cleos/" 
keyPairFile=${EOS_CLEOS_PATH}"keyPair.txt"
informationFile=${EOS_CLEOS_PATH}"information_need_for_create_account.txt"
currentLine=0
publicKey="" 
defaultWallet=""
passwordWallet=""
accountName=""
#Read wallet name + wallet password + account name need to be created
cd
while IFS= read line 
do
    let "currentLine++"
    # if [ $currentLine = 2 ]
    # then
    #     defaultWallet=$line
    # fi
    case $currentLine in
        1) accountName="${line:14}"
            ;;
        2) defaultWallet="${line:13}"
            ;;
        3) passwordWallet=${line:17}
            ;;
    esac
done < $informationFile
currentLine=0
echo "Wallet name: "
echo "${defaultWallet}"
echo "Account name: "
echo "${accountName}"
echo "Wallet password: "
echo "${passwordWallet}"
#Create key. 1 key - 1 account
cd $EOS_CLEOS_PATH
./cleos create key -f $keyPairFile
cd

#Unlock wallet
cd ~/eos/build/programs/cleos/
./cleos wallet unlock -n "${defaultWallet}" --password "${passwordWallet}" 
cd
# Read public key from the one created above
while IFS= read line 
do
    let "currentLine++"
    if [ $currentLine = 2 ]
    then
        publicKey=$line
    fi
done < "$keyPairFile"
publicKey="${publicKey:12}"
echo "Public key: "
echo $publicKey
cd $EOS_CLEOS_PATH
./cleos create account eosio $accountName $publicKey
cd
