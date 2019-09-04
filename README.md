# APC PDU AP7900

APC PDU AP7900のコントロールと情報を取得します

# 準備

* Telnetを使用可能にしておく

# ビルド

golang, make が必要です。

	$ make

# 利用方法

	$ cp ap7900.yaml.example ap7900.yaml
	$ vim ap7900.yaml
	$ ./ap7900 -l

# コマンドラインオプション

オプション | 内容
-----------|--------
-c         | 設定ファイル、省略時はカレントディレクトリの ap7900.yaml
-l         | 負荷(単位:アンペア)
-o         | 出力ポート
-s         | 状態
-on        | すぐにON
-off       | すぐにOFF
-reboot    | すぐにOFFにしてON
-don       | 遅延してON
-doff      | 遅延してOFF
-dreboot   | 遅延してOFFにしてON

# 使用例

Outlet 1をすぐにONにする

	$ ./ap7900 -o 1 -on

Outlet 1の状態

	$ ./ap7900 -o 1 -s

負荷

	$ ./ap7900 -s
