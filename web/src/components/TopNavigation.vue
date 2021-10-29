<template>
  <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
  >
    <el-menu-item index="1" @click="goHome">首页</el-menu-item>
    <el-menu-item index="2">数据集</el-menu-item>
    <el-menu-item index="3">方法</el-menu-item>
    <el-menu-item index="4" @click="goNews">新闻</el-menu-item>
    <el-menu-item index="5" @click="goAbout">关于</el-menu-item>
  </el-menu>
</template>

<script>
export default {
  name: "TopNavigation",
  methods: {
    goHome() {
      this.$router.push('/')
    },
    goNews(){
      this.$router.push('/news')
    },
    goAbout(){
      this.$router.push('/about')
    }
  }
}
</script>

<style scoped>

</style>