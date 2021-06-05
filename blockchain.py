# coding: UTF-8

class Blockchain(object):
    def __init__(self):
        self.current_transactions = []
        self.chain = []

    def new_transaction(self, sender, recipient, amount):
        """
        次に採掘されるブロックに加える新しいトランザクションを作成
        :param sender: <str> 送信者のアドレス
        :param recipient: <str> 受信者のアドレス
        :param amount: <int> 量
        :return: <int> このトランザクションを含むブロックのアドレス
        """

    def new_block(self, proof, previous_hash=None):
        """
        ブロックチェーンに新しいブロックを作る
        :param proof: <int> プルーフオブワークアルゴリズムから得られるプルーフ
        :param previous_hash: <str> 直前のブロックのハッシュ
        :return: <dict> 新しいブロック
        """

    @property
    def last_block(self):
        return self.chain[-1]

    @staticmethod
    def hash(block):
        """
        ブロックの　SHA-256　ハッシュを作る
        :param block: <dict> ブロック
        :return: <str>
        """