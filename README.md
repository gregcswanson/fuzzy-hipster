fuzzy-hipster
=============

Check and Recheck web application

Launch from Nitrious.IO
dev_appserver.py --address=0.0.0.0 --port=3000 workspace/fuzzy-hipster/
goapp serve -host=0.0.0.0 -port=3000 workspace/go/src/github.com/fuzzy-hipster/

Deploy to google app engine
appcfg.py update workspace/fuzzy-hipster/
