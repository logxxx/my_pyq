<template>
  <div>
  <el-button circle size="large" type="primary" style="position: fixed;top:30px; right:30px;" @click="addRecord"><el-icon><Plus /></el-icon></el-button>
  <el-button circle size="large" type="success" style="position: fixed;top:100px; right:30px;" @click="onSubmitRecord"><el-icon><Select /></el-icon></el-button>
  <br/>
  <el-card v-for="(elem, id) in list">
    <div>序号: {{id+1}}</div>
    <div>标题：<el-input style="width:300px" v-model="elem.title" placeholder="Please input" /></div>
    <div>素材：<el-upload
        class="upload-demo"
        drag
        multiple
        action="http://192.168.31.5:8090/backend/v1/file/upload"
        :on-success="function(res, file) {return onUploadSource(res, file, elem)}"
        :on-remove="function(file) {return onRemoveSource(file, elem)}"
        list-type="picture-card"
    ><el-icon><Plus /></el-icon></el-upload></div>
    <div>
      <div >点赞:</div>
      <div>生成<el-input-number v-model="geneLikesCount"></el-input-number>个 <el-button @click="geneLikes(elem)">√</el-button></div>
      <el-tag :color="randomColor()" style="margin-right: 10px;color:white;" v-for="elem in elem.likes">{{elem}}</el-tag>

    </div>
    <div>
      <div>回复:</div>
      <div>生成<el-input-number v-model="geneReplyCount"></el-input-number>个 <el-button @click="geneReply(elem)">√</el-button></div>
      <div v-for="r in elem.replies">
        <el-input v-model="r.name" style="width:150px;"></el-input> 说: <el-input v-model="r.content" style="width:250px;"></el-input>
      </div>
      <el-button style="marign-top:30px;" @click="onClickAddReply(elem)">+</el-button>
    </div>
  </el-card>
  </div>
</template>
<script>
import { Plus } from '@element-plus/icons-vue'
import { Select } from '@element-plus/icons-vue'
import axios from "axios";

export default {
  name: 'Mgr',
  props: {
    msg: String
  },
  components: {
    Plus,
    Select,
  },
  data() {
    return {
      geneLikesCount: 3,
      geneReplyCount: 3,
      colorList: [
        '#C0392B', '#E74C3C', '#F39C12', '#1ABC9C', '#3498DB',
        '#9B59B6', '#F1C40F', '#2ECC71', '#34495E', '#95A5A6',
        '#7F8C8D', '#F8C471', '#F89814', '#D35400', '#EA7E53',
        '#E74C3C', '#8257E6', '#2979FF', '#40E0D0', '#A5D6A7',
        '#33691E', '#CCA010', '#B85450', '#883E3E', '#728AAD',
      ],
      list: [{
        title: '',
        sources: [],
        replies: [],
        likes: [],
      }],
    }
  },
  computed: {

  },
  mounted(){

  },
  created(){

  },
  methods: {
    addRecord: function() {
      this.list.push({
        title: '',
        sources: [],
        replies: [],
      })
    },

    geneLikes(elem) {
      axios.get('http://192.168.31.5:8090/backend/v1/gene_likes?count='+this.geneLikesCount).then(resp=>{
          console.log("gene_likes resp:", resp.data)
          elem.likes = []
          resp.data.forEach(v=>{
            elem.likes.push(v)
          })
      })
    },

    geneReply(elem) {
      axios.get('http://192.168.31.5:8090/backend/v1/gene_reply?count='+this.geneReplyCount).then(resp=>{
        console.log("gene_reply resp:", resp.data)
        elem.replies = []
        resp.data.forEach(v=>{
          elem.replies.push({name:v.name, content:v.content})
        })
        console.log("elem.replies:", elem.replies)
      })
    },

    randomColor() {
      const randomIndex = Math.floor(Math.random() * this.colorList.length);
      return this.colorList[randomIndex];
    },

    onRemoveSource: (file, elem) => {
      console.log("remove file:", file)
      console.log("remove elem before:", elem)
      elem.sources.forEach((v,idx)=>{
        if(v.name===file.name){
          elem.sources.splice(idx, 1)
          return
        }
      })
      console.log("remove elem after:", elem)
    },

    onClickAddReply: (elem) => {
      elem.replies.push({})
    },

    onUploadSource: (resp, file, elem) => {
      if(resp && resp.id){
        elem.sources.push({name: file.name, id:resp.id})
      }
      if(elem.title == "") {
        var startIdx = file.name.indexOf("_")
        console.log("startIdx:", startIdx)
        if(startIdx==0){
          startIdx = file.name.indexOf("_", 1)
          console.log("mew startIdx:", startIdx)
        }
        var endIdx = file.name.indexOf("_", startIdx+1)
        if(endIdx<=0){
          endIdx = file.name.indexOf(".", startIdx+1)
        }
        console.log("endIdx:", endIdx)
        elem.title = file.name.substring(startIdx+1, endIdx)
      }
      console.log("upload resp:", resp)
      console.log("upload file:", file)
      console.log("upload elem:", elem)
    },

    onSubmitRecord: function(){
      axios.post('http://192.168.31.5:8090/backend/v1/apply_tmpl', {data:this.list})
    },
  }
}
</script>