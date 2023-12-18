from bitcoinutils.keys import PrivateKey, P2pkhAddress
from bitcoinutils.setup import setup
from bitcoinutils.transactions import Transaction, TxInput, TxOutput
from bitcoinutils.script import Script
from bitcoinutils.utils import to_satoshis

if __name__ == "__main__":
    private_key = PrivateKey()
    public_key = private_key.get_public_key()
    address = P2pkhAddress(public_key.get_address().to_string())
    print("Private key: " + private_key.to_wif())
    print("Public key: " + public_key.to_hex())
    print("Address: " + address.to_string())