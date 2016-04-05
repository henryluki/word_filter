# coding: utf-8
from flask import Flask, request, jsonify, render_template
from classify import MyGrocery
import logging, sys

app = Flask(__name__)
# log
app.logger.addHandler(logging.StreamHandler(sys.stdout))
app.logger.setLevel(logging.DEBUG)

grocery = MyGrocery("SVM")
grocery.predict("你好")

def request_params(request):
  if len(request.form) != 0:
    return request.form
  if len(request.json) != 0:
    return request.json

@app.route('/')
def index():
  return render_template('index.html')

@app.route('/classify', methods=['POST'])
def classify():
  text = request_params(request)['text'].strip(' ')
  label = grocery.predict(text)
  app.logger.info("[INFO] label: "+ label + " text: " + text)
  return jsonify({ "label": label, "text": text})

if __name__ == '__main__':
  app.run(debug=True, port=8006, host='0.0.0.0')