$(document).ready(function() {
  var postData = function (data, success_cb){
    $.ajax({
      type:'post',
      url: 'classify',
      data: JSON.stringify(data),
      dataType:'json',
      beforeSend:function(xhr){
        xhr.setRequestHeader('Content-Type', 'application/json' )
      },
      success:function(res){
        success_cb(res);
      },
      error:function(){
        console.log('failed!');
      }
    });
  }

  var handleClick = function(){
    var text = $("#input").val();
    var predictArea = $("#predict");
    if (text != "") {
      var data = {
        text: text
      };
      postData(data, function(data){
        var polyfill = []
        polyfill.push("<p>文本：<span class='sample-text'>", data.text, "</span> 类型: <span class='sample-label'>", data.label, "</span></p>")
        predictArea.append(polyfill.join(""))
      });

    }else{
      alert("输入不能为空！")
    }
  };

  $("#button").bind("click", handleClick)
});