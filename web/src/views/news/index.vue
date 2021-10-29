<template>
  <ul>
    <li v-for="item in news" :key="item.id">
      <p>{{item}}</p>
    </li>
  </ul>
</template>

<script>
import axios from "axios";

export default {
  name: "news",
  created(){
    axios.get("http://localhost:8081/news/getAllNews").then(res => {
      this.news = res.data
    })
  },
  data(){
    return {
      news:[]
    }
  },
  methods:{

  }
}
</script>

<style scoped>
ul,li{ padding:0;margin:0;list-style:none}
</style>