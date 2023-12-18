from bitcoinutils.keys import PrivateKey, P2pkhAddress, P2shAddress
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
    address1 = P2pkhAddress(public_key.get_address().to_string())
    print("Address: " + address1.to_string())

    # Wallet 2:
    private_key2 = PrivateKey.from_wif(os.getenv("PRIVATE_KEY2"))
    public_key2 = private_key2.get_public_key()
    address2 = P2pkhAddress(public_key2.get_address().to_string())
    print("Address2: " + address2.to_string())
    out_address = P2pkhAddress(os.getenv("OUT_ADDRESS"))

    # P2SH multisig transaction example
    tx_in = TxInput(os.getenv("UTXO_SH"), 1)
    redeem_script = Script(['OP_2', public_key.to_hex(), public_key2.to_hex(), 'OP_2', 'OP_CHECKMULTISIG'])
    print("Redeem script: " + redeem_script.to_hex())
    p2sh_address = P2shAddress.from_script(redeem_script)
    print("P2SH address: " + p2sh_address.to_string())

    tx_out = TxOutput(to_satoshis(0.001), out_address.to_script_pub_key())
    tx = Transaction([tx_in], [tx_out],has_segwit=True)
    print("\nRaw unsigned transaction:\n" + tx.serialize())

    sig = private_key.sign_input(tx, 0, redeem_script)
    sig2 = private_key2.sign_input(tx, 0, redeem_script)
    tx_in.script_sig = Script(['OP_0',sig, sig2, redeem_script.to_hex()])
    print("\nRaw signed transaction:\n" + tx.serialize())

    print("\nTransaction is fully signed and ready for broadcast")
    print(pushtx(tx.serialize(), coin_symbol='btc-testnet', api_key=os.getenv("BC_TOKEN")))
    print("Transaction broadcasted!")

    
