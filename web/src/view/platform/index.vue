<template>
  <div class="iframestyleset">
    <iframe name="iframeMap" id="iframeMapViewComponent" v-bind:src="getPageUrl"
            width="100%" :height="windowHeight" ref="iframeDom"
            @load="loadJump"
    ></iframe>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
const path = import.meta.env.VITE_BASE_PATH+":"+import.meta.env.VITE_SERVER_PORT

export default {
  name: 'Platform',
  data() {
    return {
      baseurl: path,
      windowWidth: document.documentElement.clientWidth, //实时屏幕宽度
      windowHeight: document.documentElement.clientHeight, //实时屏幕高度
      getPageUrl: 'public/job.html',
    }
  },
  created() {
    window['quit'] = (data) => {
      this.quit()
    }

  },
  computed: {
    ...mapGetters('user', ['token']),
  },
  methods: {
    ...mapActions('user', ['LoginOut', 'GetUserInfo']),
    loadJump() {

      let allApps = [{
        name: '网信监测',
        icon: './images/wxjc.png',
        url: this.baseurl+'/Jq/getJunquanJump?x-token='+this.token+"&"+new Date().getTime(),
      }, {
        name: '网信考核',
        icon: './images/wxkh.png',
        url: '',
      }, {
        name: '网信考核',
        icon: './images/wxkh.png',
        url: '',
      }, {
        name: '舆情报告',
        icon: './images/wxkh.png',
        url: '',
      }]
      window.frames['iframeMap'].setApps(allApps)

    },
    quit() {
      console.log('quit')
      this.LoginOut()
    },

  },
}
</script>

<style scoped>
</style>
