<template>
  <div class="app-bar">
    <div class="left-links">
      <router-link
          class="app-bar-item"
          v-for="link in links"
          :key="link.name"
          :to="link.to">{{ link.name }}</router-link>
    </div>
    <div class="right-links ">
      <a class="app-bar-item" href="#">LOGIN</a>
    </div>
  </div>
</template>
<script>
export default
{
  data() {
    return {
      links: [
        {
          name: "Предметы",
          to: { name: "MySubjects" }
        }
      ]
    };
  }};
</script>
<style scoped>
.app-bar {
  height: 5vh;
  width: 100%;
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  background-color: #02060e;
  overflow: hidden;
  display: block;
  justify-content: space-between;
}
.left-links {
  padding-left: 10%;
}
.left-links a {
  float: left;
  color: rgba(255, 255, 255, 0.65);
}

a:hover {
  opacity: 0.5;
}

a {
  float: left;
  display: block;
  color: #ffffff;
  text-align: center;
  padding: 14px 48px;
  font-weight: bold;
  text-decoration: none;
  font-size: 17px;
}
.right-links {
  float: right;
}
</style>