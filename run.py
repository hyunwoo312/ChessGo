from lib.webpy import app
import sys

if __name__=="__main__":
    if sys.argv[1] == "--launch":
        app.run(debug=True)