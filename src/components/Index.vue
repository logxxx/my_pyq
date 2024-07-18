<template>
  <van-config-provider theme="dark"></van-config-provider>
  <div id="main_page" style="overflow:hidden;" class="w-full h-full bg-[#f1f5f9] dark:bg-slate-800 rounded-md dark:text-[#C0BEBF]">
    <div class="h-full">
      <div class="main lg:w-[567px] mx-auto shadow-2xl bg-[#181818]">
        <!--HeaderImg start -->
        <div class="header relative mb-8 ">

          <img class="header-img w-full h-auto" :src="getImgUrl(profile.cover)" alt="" />

          <div class="absolute right-4 bottom-[-40px]">
            <div class="userinfo flex flex-col">
              <div class="flex flex-row gap-2 justify-end">
                <div class="username text-lg font-bold text-white">{{ profile.nickname }}</div>
                <img :src="getImgUrl(profile.avatarUrl)" class="avatar w-[70px] h-[70px] rounded-xl"/>
              </div>
            </div>
          </div>

        </div>

        <!--HeaderImg end-->


        <div class="mt-12">

        <div class="content flex flex-col divide-y divide-[#C0BEBF]/10">
          <FriendsMemo :memo="memo" v-for="(memo, index) in memos" :index="index" :key="memo.id" :show-more="true"/>
        </div>

      </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import FriendsMemo from '@/components/FriendsMemo.vue'


export default {
  name: 'Index',
  components: {
    FriendsMemo,
  },
  created(){
    this.getProfile()
    this.getMemos()
  },
  data(){
    return {
      memos: [],
      profile: {},
    }
  },
  methods: {

    getProfile: function() {
      axios.get('http://192.168.50.250:8090/backend/v1/profile').then(resp => {
        this.profile = resp.data
        console.log("getProfile resp:", this.profile)
      })
    },

    getImgUrl: function(input){
      return "http://192.168.50.250:8090/backend/v1/file?id=" + input;
    },

    getMemos: function() {
      axios.get('http://192.168.50.250:8090/backend/v1/memos').then(resp => {
        resp.data.memos.forEach(v=>{
          this.memos.push(v)
        })
        console.log("getMemos resp:", this.memos)
      })
    },



  }

}
</script>

<style>
#main_page {
  overflow-y: scroll;
  scrollbar-width: none;
}

#main_page::-webkit-scrollbar {
  display: none; /* 对于 Chrome, Safari 和 Opera 的写法 */
}

</style>