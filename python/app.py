from flask import Flask

app = Flask(__name__)

@app.route('/')
def hey():
    return '<p style="text-align: center; padding-top: 10%;">Aren\'t you a good k8s: <b style="font-size: 2rem;">v1 from flask</b></p>'

if __name__ == '__main__':
    app.run()
