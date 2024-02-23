from flask import Flask
app = Flask(__name__)

@app.route('/traffic_test')
def some_task():
    for i in range(10000):
        a = i ** i
    return "Traffic Test Request!"

if __name__ == '__main__':
    app.run(host='0.0.0.0')
