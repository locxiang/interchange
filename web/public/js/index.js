function setTime() {
  var now = dayjs()
  $('#time').text(now.format('HH:mm:ss'))
  $('#date').text(now.format('YYYY-MM-DD'))
  $('#week').text(now.format('dddd'))
}

setTime()
setInterval(function() {
  setTime()
}, 1000)

function setApps(allApps) {
  allApps.forEach(app => {
    $('.apps').append(`
<a target="_blank" href="${app.url}" >
    <div class="app">
        <div class="app-icon">
           <img src="${app.icon}" alt="">
        </div>
        <div class="app-name"> ${app.name}</div>
    </div>
    </a>
    `)
  })
}

//
// // 应用点击事件
// $('.app').on('click', function (e){
//   var $dom = $(e.target)
//   var clickApp = allApps.find(v => v.name === $dom.text().replace(/\s/g, ''));
//   console.log(clickApp)
//   alert('点击了应用：' + clickApp.name)
// })

// 切换个人中心显示
let showApps = true
$('#personCenter').on('click', function() {
  if (showApps) {
    $('#personal').fadeIn()
    $('#apps').fadeOut()
    showApps = false
  } else {
    $('#personal').fadeOut()
    $('#apps').fadeIn()
    showApps = true
  }
})

//个人设置点击事件
function f() {
  document.getElementsByClassName('dropdown')[0].classList.toggle('down')
  document.getElementsByClassName('arrow')[0].classList.toggle('gone')
  if (document.getElementsByClassName('dropdown')[0].classList.contains('down')) {
    setTimeout(function() {
      document.getElementsByClassName('dropdown')[0].style.overflow = 'visible'
    }, 500)
  } else {
    document.getElementsByClassName('dropdown')[0].style.overflow = 'hidden'
  }
}

// 修改名称
$('#editName').on('click', function() {
  $('#editDiv').hide()
  $('#formDiv').show()
})
// 保存修改名称
$('#saveName').on('click', function() {
  $('#editDiv').show()
  $('#formDiv').hide()
})
// 取消修改名称
$('#cancelSave').on('click', function() {
  $('#editDiv').show()
  $('#formDiv').hide()
})
