# coding: utf-8
from flask import Flask, request, jsonify
from classify import MyGrocery

app = Flask(__name__)

grocery = MyGrocery("SVM")
grocery.predict("你好")

@app.route('/')
def index():
  return 'hello world'

@app.route('/classify', methods=['POST'])
def classify():
  text = request.form['text']
  label = grocery.predict(text)
  return jsonify({ "label": label, "text": text})

if __name__ == '__main__':
  app.run(debug=True, port=8006, host='0.0.0.0')