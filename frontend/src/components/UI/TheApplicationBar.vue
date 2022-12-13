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
      <a class="app-bar-item" href="#" v-if="!loggedIn" @click.prevent="login">LOGIN</a>
      <a class="app-bar-item" href="#" v-if="loggedIn" @click.prevent="logout">LOGOUT</a>
    </div>
  </div>
</template>
<script>
import { mapGetters } from "vuex";
import { mapActions } from "vuex";

export default
{
  data() {
    return {
      links: [
        {
          name: "Предметы",
          to: { name: "MySubjects" }
        },
      ]
    };
  },
  methods: {
    ...mapActions({
      login: "auth/login",
      logout: "auth/logout"
    })
  },
  computed: {
    ...mapGetters({ loggedIn: "auth/isLoggedIn" })
  }
};
</script>
<style scoped>
.app-bar {
  height: 5vh;
  width: 100%;
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  background-color: #ffffff;
  box-shadow: 0 5px 0 rgba(210, 210, 210, 1);
  overflow: hidden;
  display: block;
  justify-content: space-between;
}
.left-links {
  padding-left: 10%;
}
.left-links a {
  float: left;
  color: rgba(5, 5, 5, 0.65);
}

a:hover {
  opacity: 0.5;
}

a {
  float: left;
  display: block;
  color: #000000;
  text-align: center;
  padding: 14px 0;
  font-weight: bold;
  text-decoration: none;
  font-size: 17px;
}
.right-links {
  float: right;
  padding-right: 10%;
}
</style>