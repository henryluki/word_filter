# coding: utf-8
from flask import Flask, request, jsonify, render_template
from werkzeug import secure_filename
from classify import MyGrocery
import logging, sys, os

app = Flask(__name__)
# upload config
UPLOAD_FOLDER = '../../data/uploads'
ALLOWED_EXTENSIONS = set(['csv'])
app.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER

# log
app.logger.addHandler(logging.StreamHandler(sys.stdout))
app.logger.setLevel(logging.DEBUG)

# trained model init
grocery = MyGrocery("SVM")
grocery.predict("你好")

# utils
def request_params(request):
  if len(request.form) != 0:
    return request.form
  if len(request.json) != 0:
    return request.json

def allowed_file(filename):
  return '.' in filename and \
    filename.rsplit('.', 1)[1] in ALLOWED_EXTENSIONS

# routes
@app.route('/')
def index():
  return render_template('index.html')

@app.route('/upload', methods=['post'])
def upload():
  if request.method == 'POST':
    file = request.files['file']
    if file and allowed_file(file.filename):
      filename = secure_filename(file.filename)
      file.save(os.path.join(app.config['UPLOAD_FOLDER'], filename))
      res = { "status": 200, "filename": filename }
    else:
      res = { "status": 500 }
  return jsonify(res)


@app.route('/action', methods=['POST'])
def action():
  params = request_params(request)
  action_type = params['type']
  filename = params['filename']
  src = UPLOAD_FOLDER + '/' + filename
  if action_type == "train":
    grocery.train_and_save(src)
    return jsonify({ "type" : 'train'})
  elif action_type == "test":
    result = grocery.test(src)
    return result

@app.route('/classify', methods=['POST'])
def classify():
  text = request_params(request)['text'].strip(' ')
  label = grocery.predict(text)
  app.logger.info("[INFO] label: "+ label + " text: " + text)
  return jsonify({ "label": label, "text": text})

if __name__ == '__main__':
  app.run(debug=True, port=8006, host='0.0.0.0')