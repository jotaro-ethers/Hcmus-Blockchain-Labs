from bitcoinutils.keys import PrivateKey, P2pkhAddress
from bitcoinutils.setup import setup
from bitcoinutils.transactions import Transaction, TxInput, TxOutput
from bitcoinutils.script import Script
from bitcoinutils.utils import to_satoshis
from blockcypher import pushtx
from dotenv import load_dotenv
import os
load_dotenv()

if __name__ == "__main__":
    setup("testnet")
    # Wallet 1:
    private_key = PrivateKey.from_wif(os.getenv("PRIVATE_KEY"))
    public_key = private_key.get_public_key()
    address = P2pkhAddress(public_key.get_address().to_string())
    print("Address: " + address.to_string())
    out_address = P2pkhAddress(os.getenv("OUT_ADDRESS"))
    
    #P2PKH transaction example 
    tx_in = TxInput(os.getenv("UTXO_KH"), 1)
    tx_out = TxOutput(to_satoshis(0.00001), Script(['OP_DUP', 'OP_HASH160',out_address.to_hash160(), 'OP_EQUALVERIFY', 'OP_CHECKSIG']))
    change_txout = TxOutput(to_satoshis(0.001), Script(['OP_DUP', 'OP_HASH160', address.to_hash160(), 'OP_EQUALVERIFY', 'OP_CHECKSIG']))
    tx = Transaction([tx_in], [tx_out, change_txout])
    print("\nRaw unsigned transaction:\n" + tx.serialize())

    sig = private_key.sign_input(tx, 0, Script(['OP_DUP', 'OP_HASH160', address.to_hash160(), 'OP_EQUALVERIFY', 'OP_CHECKSIG']))
    tx_in.script_sig = Script([sig, public_key.to_hex()])

    print("\nRaw signed transaction:\n" + tx.serialize())
    print(pushtx(tx.serialize(), coin_symbol='btc-testnet', api_key=os.getenv("BC_TOKEN")))

    