<template>
  <div class="app-bar">
    <div class="left-links">
      <router-link
          class="app-bar-item"
          v-for="link in activeLinks"
          :key="link.name"
          :to="link.to">{{ link.name }}</router-link>
    </div>
    <div class="right-links ">
      <router-link class="app-bar-item" href="#" v-if="!loggedIn" @click.prevent :to="{ name: 'Login' }" LOGIN>LOGIN</router-link>
      <a class="app-bar-item" href="#" v-if="loggedIn" @click.prevent="logoutButtonClicked">LOGOUT</a>
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
          visibleIfLoggedOut: false,
          name: "Предметы",
          to: { name: "StudentSubjects" }
        },
      ]
    };
  },
  methods: {
    ...mapActions({
      login: "auth/login",
      logout: "auth/logout"
    }),
    logoutButtonClicked() {
      this.logout().then(() => {
        this.$router.push({ name: "Login" });
      });
    }
  },
  computed: {
    ...mapGetters({ loggedIn: "auth/isLoggedIn" }),
    activeLinks() {
      return this.links.filter(
          link => link.visibleIfLoggedOut || this.loggedIn
      )
    }
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