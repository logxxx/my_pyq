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

      <video class="w-2/3 my-2 rounded" :src="getFile(memo.video)" autoplay="autoplay" controls v-if="memo.video"></video>

      <div v-if="memo.images.length">
        <div class="grid my-1 gap-0.5" :style="getImgGridStyle(memo)" :options="{ Carousel: { infinite: false } }">
          <img loading="lazy" :src="getFile(img)" class="cursor-zoom-in"
               :class="memo.images.length === 1 ? 'full-cover-image-single' : 'full-cover-image-mult'"
               v-for="(img, index) in memo.images" :key="index" />
        </div>
      </div>

      <div style="overflow: hidden;" class="toolbar relative flex flex-row justify-between select-none my-3 items-center">

        <div class="flex-1 text-gray text-xs text-[#9DA4B0] " >{{
            memo.show_time
          }}</div>

        <!--三个点-->
        <div style="z-index:2; background-color: #191b1f;height:40px;width:30px;position:absolute;right:0px;">

        </div>
        <div @click="showToolbar = !showToolbar"
             style="z-index:3;background-color: rgb(44 44 44)" class="toolbar-icon px-2 py-1 bg-[#f7f7f7] dark:bg-slate-700 hover:bg-[#dedede] cursor-pointer rounded flex items-center justify-center">
          <img src="@/assets/img/dian.svg" class="w-3 h-3" />
        </div>

        <!--弹出:赞或评论-->
        <Transition  name="toolbarmove2">
        <div style="z-index:1;" class="absolute top-[-8px] right-[32px] bg-[#4c4c4c] rounded text-white p-2 px-4" v-if="showToolbar"
             ref="toolbarRef">

          <div class="flex flex-row gap-4">
            <div class="flex flex-row gap-0.5 cursor-pointer items-center" @click="like">
              <Heart style="margin-left:10px" :size=14 v-if="!memo.is_like" />
              <Heart style="margin-left:10px;color:red;fill: red;" :size=14 v-else />
              <div @click="memo.is_like = !memo.is_like;showToolbar=false" style="margin-right:10px">{{ memo.is_like ? '取消' : '赞' }}</div>
            </div>
            <span class="bg-[#6b7280] h-[20px] w-[1px]"></span>

            <div class="flex flex-row gap-0.5 cursor-pointer items-center"
                 @click="momentsShowCommentInput = !momentsShowCommentInput; showUserCommentArray = []; showToolbar = false">
              <MessageSquare :size=14 />
              <div style="min-width:10px;white-space: nowrap;margin-right: 5px;">评论</div>
            </div>
          </div>

        </div>
        </Transition>


      </div>

    </div>
    </div>
  </div>
</template>
<script>
import {Heart, HeartCrack, MessageSquare} from "lucide-vue-next";
export default {
  name: 'FriendsMemo',
  components:{
    Heart,
    HeartCrack,
    MessageSquare,
  },
  props: {
    memo: Object,
  },
  data() {
    return {
      showToolbar: false,
      momentsShowCommentInput: false,
      showUserCommentArray: [],
    }
  },
  mounted(){

  },
  created(){

  },
  methods: {

    getToolbarStyle() {

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


/* 定义动画 */
@keyframes slideFadeIn {
  0% {
    max-height: 36px;
    width: 0%;
  }
  100% {
    max-height: 36px;
    width: 150px; /* 假设最大高度是100px */
  }
}

@keyframes slideFadeIn2 {
  0% {
    transform: translateX(100%);
    opacity: 0;
  }
  100% {
    transform: translateX(0);
    opacity: 100%;

  }
}

@keyframes slideFadeOut2 {
  0% {
    transform: translateX(0%);
    opacity: 100%;

  }
  10% {
    transform: translateX(0%);
    opacity: 100%;
  }
  100% {
    transform: translateX(100%);
    opacity: 0%;

  }
}

@keyframes slideFadeOut {
  0% {
    max-height: 36px;
  }
  100% {
    max-height: 36px;
    width: 0; /* 假设最大高度是100px */
  }
}

/* 应用动画 */
.popup {
  animation-name: slideFadeIn;
  animation-duration: 0.5s; /* 动画持续时间 */
  animation-fill-mode: forwards; /* 动画结束后保持最后一帧的样式 */
  overflow: hidden; /* 隐藏超出容器的内容 */
}


p, span {
  letter-spacing: 1px;
}



.toolbarmove2-enter-active {
  animation-name: slideFadeIn2;
  animation-duration: 0.5s; /* 动画持续时间 */
  animation-fill-mode: forwards; /* 动画结束后保持最后一帧的样式 */
  overflow: hidden; /* 隐藏超出容器的内容 */
}

.toolbarmove2-leave-active {
  animation-name: slideFadeOut2;
  animation-duration: 0.5s; /* 动画持续时间 */
  animation-fill-mode: forwards; /* 动画结束后保持最后一帧的样式 */
  overflow: hidden; /* 隐藏超出容器的内容 */
}

.toolbarmove-enter {
  width: 0;
}

.toolbarmove-enter-active {
  animation-name: slideFadeIn;
  animation-duration: 0.3s; /* 动画持续时间 */
  animation-fill-mode: forwards; /* 动画结束后保持最后一帧的样式 */
  overflow: hidden; /* 隐藏超出容器的内容 */
}

.toolbarmove-enter-to {
  width: 150px
}

.toolbarmove-leave {

}

.toolbarmove-leave-active {
  animation-name: slideFadeOut;
  animation-duration: 0.3s; /* 动画持续时间 */
  animation-fill-mode: forwards; /* 动画结束后保持最后一帧的样式 */
  overflow: hidden; /* 隐藏超出容器的内容 */
}

.toolbarmove-leave-to {

}







</style>
