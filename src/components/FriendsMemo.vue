<template>
  <div>
    <!--<div class="memo flex flex-row gap-2 sm:gap-4 text-sm border-x-0 pt-2 p-2 sm:p-4">-->
    <div class="memo flex flex-row gap-2 sm:gap-4 text-sm border-x-0 pt-2 p-2 sm:p-4">

      <img :src="getFile(memo.user.avatarUrl)" class="avatar w-10 h-10 rounded "/>

      <div>

        <div class="flex flex-col gap-.5 flex-1">
          <div class="">
            <div style="font-size: 1.2rem; margin-top: -5px; color: #7e90a8" class="username text-lg dark:text-white cursor-pointer">
              {{memo.user.nickname }}
            </div>
            <div style="font-size: 1.2rem" class="memo-content text-lg friend-md mome-container text-wrap break-all" ref="el" v-if="memo.content">{{memo.content}}</div>
        </div>
      </div>

        <div :id="memo.video" ref="myVideoContainer" class="orig-video">
          <video ref="myVideo" @click="videoFullScreen(memo.video)" :src="getFile(memo.video)" autoplay="autoplay" v-if="memo.video"></video>
        </div>

      <div v-if="memo.images && memo.images.length>0">
        <div class="grid my-1 gap-0.5" :style="getImgGridStyle(memo)" :options="{ Carousel: { infinite: false } }">
          <img @click="showImgPreview=true" loading="lazy" :src="getFile(img)" class="cursor-zoom-in"
               :class="memo.images.length === 1 ? 'full-cover-image-single' : 'full-cover-image-mult'"
               v-for="(img, index) in memo.images" :key="index" />
        </div>
        <van-image-preview v-model:show="showImgPreview" :images="getImgs(memo.images)">
        </van-image-preview>

      </div>


      <div style="color: #7e90a8;font-size: 1.1rem">{{memo.place}}</div>

      <div style="" class="toolbar relative flex flex-row justify-between select-none my-3 items-center">

        <div class="flex-1 text-gray text-m text-[#9DA4B0] " >{{
            memo.show_time
          }}</div>

        <!--三个点-->
        <div style="right: -8px; padding-right: 8px; z-index:2; background-color: #191b1f;height:40px;width:40px;position:absolute;">
        </div>
        <div @click="showToolbar = !showToolbar"
             style="z-index:3;background-color: rgb(44 44 44)" class="toolbar-icon px-2 py-1 bg-[#f7f7f7] dark:bg-slate-700 hover:bg-[#dedede] cursor-pointer rounded flex items-center justify-center">
          <img src="@/assets/img/dian.svg" class="w-3 h-3" />
        </div>

        <!--弹出:赞或评论-->
        <Transition  name="toolbarmove">
        <div style="z-index:1;" class="absolute top-[-8px] right-[32px] bg-[#4c4c4c] rounded text-white p-2 px-4" v-if="showToolbar"
             ref="toolbarRef">

          <div class="flex flex-row gap-4">
            <div class="flex flex-row gap-0.5 cursor-pointer items-center" @click="like">
              <Heart style="margin-left:10px" :size=14 v-if="!memo.is_like" />
              <Heart style="margin-left:10px;color:red;fill: red;" :size=14 v-else />
              <div @click="memo.is_like = !memo.is_like;showToolbar=false;memo.likes?memo.likes.push(profile.nickname):memo.likes=[profile.nickname]" style="margin-right:10px">{{ memo.is_like ? '取消' : '赞' }}</div>
            </div>
            <span class="bg-[#6b7280] h-[20px] w-[1px]"></span>

            <div class="flex flex-row gap-0.5 cursor-pointer items-center"
                 @click="momentsShowCommentInput = !momentsShowCommentInput; showUserCommentArray = []; showToolbar = false">
              <MessageSquare :size=14 />
              <div @click="onClickComment" style="min-width:10px;white-space: nowrap;margin-right: 5px;">评论</div>
            </div>
          </div>

        </div>
        </Transition>

      </div>

        <div v-if="memo.likes" style="display: flex;margin-left: 10px;;margin-bottom:10px;">
          <Heart size=17 style="position:absolute;color:#7e90a8"/><div style="color: #7e90a8;font-size: 1.1rem; margin-left: 5px;">&nbsp;&nbsp;&nbsp;{{(memo.likes?memo.likes:[]).join('，')}}</div>
        </div>

        <div style="color: rgb(192,190,191); margin-left: 10px;margin-right: 10px; margin-bottom: 5px;" v-for="(comment, idx) in memo.comments">
          <div v-if="comment.reply_to"><span style="color: #7e90a8;font-size: 1.1rem">{{comment.username}}</span><span style="color: rgb(192,190,191);margin-left: 5px; margin-right:5px;font-size: 1.1rem">回复</span><span style="color: #7e90a8;font-size: 1.1rem">{{comment.reply_to}}</span>:<span style="color: rgb(192,190,191);margin-left: 5px; margin-right: 15px;font-size: 1.1rem">{{comment.content}}</span></div>
          <div v-else><span style="color: #7e90a8;font-size: 1.1rem">{{comment.username}}</span>:<span style="color: rgb(192,190,191);margin-left: 5px; margin-right: 15px;font-size: 1.1rem">{{comment.content}}</span></div>
        </div>


    </div>
    </div>

    <div v-if="showInput"  style="display: flex;">
      <van-field :id="memo.id" autocomplete="off" style="width:70%" v-model="inputContent" />
      <Laugh style="color:#f0f0f0;margin-left:10px;margin-right:10px;" :size=30 />
      <van-button  @click="onClickSend" style="font-size:15px;width:20%;margin-right: 10px">发送</van-button>
    </div>

  </div>
</template>
<script>
import {Heart, HeartCrack, MessageSquare, Laugh} from "lucide-vue-next";
export default {
  name: 'FriendsMemo',
  components:{
    Heart,
    HeartCrack,
    MessageSquare,
    Laugh,
  },
  props: {
    memo: Object,
    profile: Object,
  },
  data() {
    return {
      inputContent: '',
      showInput: false,
      showImgPreview: false,
      showToolbar: false,
      momentsShowCommentInput: false,
      showUserCommentArray: [],
      isVideoFullScreen: false,
    }
  },
  mounted(){

  },
  created(){

  },
  methods: {

    onClickComment() {

      var that = this
      this.showInput=true;
      setTimeout(function(){
        console.log("this.memo:", that.memo)
        var input = document.getElementById(that.memo.id)
        console.log("onClickComment input:", input)
        if(input){
          input.focus()
        }
      }, 200)

    },

    videoFullScreen(videoID) {

      const video = this.$refs.myVideo;

      const videoContainer = this.$refs.myVideoContainer

      if(!this.isVideoFullScreen){
        this.isVideoFullScreen = true

        var style = window.getComputedStyle(videoContainer);
        videoContainer.style.cssText = style.cssText;

        // 移除原有video容器的样式并全屏显示
        videoContainer.classList.replace('orig-video','fullscreen-video');
        videoContainer.removeAttribute('controls');

        video.play()
      } else {
        this.isVideoFullScreen = false
        var style = window.getComputedStyle(videoContainer);
        videoContainer.style.cssText = style.cssText;

        // 移除原有video容器的样式并全屏显示
        videoContainer.classList.replace('fullscreen-video', 'orig-video');
        //video.pause()
      }



      /*
      const video = this.$refs.myVideo;
      console.log("videoFullScreen video:", video)
      if (!document.fullscreenElement) {
        video.requestFullscreen().catch(err => {
          console.error(`Error attempting to enable fullscreen mode: ${err.message}`);
        });
      } else {
        if (document.exitFullscreen) {
          document.exitFullscreen();
        }
      }

       */
    },

    onClickSend() {
      console.log("get profile:", this.profile)
      this.showInput = false
      this.memo.comments.push({username:this.profile.nickname, content: this.inputContent})
      this.inputContent = ''
    },

    getImgs: function(rawImgs) {
      if(!rawImgs){
        return
      }
      var resp = []
      rawImgs.forEach(v=>{
        resp.push(this.getFile(v))
      })
      return resp
    },

    getFile: function(input){
      return "http://192.168.50.47:8090/backend/v1/file?id=" + input;
    },

    getImgGridStyle: function(memo){
      {
        let style = 'align-items: start;'; // 确保内容顶部对齐
        switch (memo.images.length) {
          case 1:
            style += 'grid-template-columns: 1fr;';
            break;
          case 2:
            style += 'grid-template-columns: 1fr 1fr 1fr ; aspect-ratio: 3 / 1;';
            break;
          case 3:
            style += 'grid-template-columns: 1fr 1fr 1fr; aspect-ratio: 3 / 1;';
            break;
          case 4:
            style += 'grid-template-columns: 1fr 1fr; aspect-ratio: 1;';
            style += "width: 65%;"
            break;
          default:
            style += 'grid-template-columns: 1fr 1fr 1fr; aspect-ratio: 3 / 1;';
        }
        return style;
      }
    }
  }
}
</script>

<style>

.memo {

}

.memo-content {
  color: rgb(192,190,191)
}


.full-cover-image-mult {
  object-fit: cover;
  object-position: center;
  width: 100%;
  aspect-ratio: 1 / 1;
  border: transparent 1px solid;
}

.full-cover-image-single {
  object-fit: cover;
  object-position: center;
  max-height: 300px;
  height: auto;
  width: auto;
  border: transparent 1px solid;
}

@keyframes slideFadeIn {
  0% {
    transform: translateX(100%);
  }
  100% {
    transform: translateX(0);
  }
}

@keyframes slideFadeOut {
  0% {
    transform: translateX(0%);
  }
  100% {
    transform: translateX(100%);
  }
}




p, span {
  letter-spacing: 1px;
}


.toolbarmove-enter-active {
  animation-name: slideFadeIn;
  animation-duration: 0.3s; /* 动画持续时间 */
}

.toolbarmove-leave-active {
  animation-name: slideFadeOut;
  animation-duration: 0.3s; /* 动画持续时间 */
}

.van-image-preview__index{
  display:none;
}

:root {
  --van-white: black;
  --van-button-default-background: red;
}

.van-button--default {
  background-color: #05c160 !important;
  border: none !important;
  height: 35px !important;
}

input2 {
  caret-color: transparent;
}

.van-cell {
  padding: 0 0 !important;
  padding-left: 20px !important;
  font-size: 17px !important;

}

.van-field__control {
  color: rgb(192,190,191) !important;
}

.fullscreen-video {
  display: flex;
  align-items: center;
  position: fixed;
  place-items: center;
  top: 0;
  left: 0;
  width: 100% !important;
  height: 100% !important;
  z-index: 9999;
  margin-top: 0 !important;
  background-color: black;
}

.orig-video {
  max-width: 60%;
  max-height:250px;
  overflow:hidden
}
</style>

