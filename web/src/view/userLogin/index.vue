<template>
  <div class="iframestyleset">
    <iframe name="iframeMap" id="iframeMapViewComponent" v-bind:src="getPageUrl"
            width="100%" :height="windowHeight" ref="iframeDom"
            @load="loadCaptcha"
    ></iframe>
  </div>
</template>

<script>
/* eslint-disable */

import { captcha } from '@/api/user'
import { mapActions } from 'vuex'
import router from '@/router/index'

export default {
  name: 'UserLogin',
  data() {
    return {
      windowWidth: document.documentElement.clientWidth, //实时屏幕宽度
      windowHeight: document.documentElement.clientHeight, //实时屏幕高度
      getPageUrl: 'public/login.html',
      captchaId: '',
    }
  },
  created() {
    window['loginVerify'] = () => {
      this.loadCaptcha()
    }
    window['login'] = (data) => {
      this.login(data)
    }

  },
  methods: {
    ...mapActions('user', ['LoginIn']),

    loadCaptcha() {
      captcha({}).then((ele) => {
        this.captchaId = ele.data.captchaId
        console.log('loadCaptchaimg',ele.data.picPath)
        window.frames['iframeMap'].loadCaptchaimg(ele.data.picPath)
      })
    },
    async login({ user, pass, captcha }) {
      let f = {
        username: user,
        password: pass,
        captcha: captcha,
        captchaId: this.captchaId,
      }
      let flag = await this.LoginIn(f)
      if (!flag) {
        this.loadCaptcha()
      }else{
        router.push({ name: "platform" })
      }
    }
  }
}
</script>

<style lang="scss">


</style>
