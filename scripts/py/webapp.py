# coding: utf-8
from flask import Flask, request, jsonify
from classify import MyGrocery
import logging, sys

app = Flask(__name__)
# log
app.logger.addHandler(logging.StreamHandler(sys.stdout))
app.logger.setLevel(logging.DEBUG)

grocery = MyGrocery("SVM")
grocery.predict("你好")

@app.route('/')
def index():
  return 'hello world'

@app.route('/classify', methods=['POST'])
def classify():
  text = request.form['text']
  label = grocery.predict(text)
  app.logger.info("[INFO] label: "+ label + " text: " + text)
  return jsonify({ "label": label, "text": text})

if __name__ == '__main__':
  app.run(debug=True, port=8006, host='0.0.0.0')