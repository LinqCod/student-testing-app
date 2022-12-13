<template>
  <div class="subject-list">
    <single-subject v-for="subject in subjects" :key="subject.id" :subject="subject">
    </single-subject>
  </div>
</template>
<script>
import SingleSubject from "./SingleSubject.vue";
export default {
  components: { SingleSubject },
  props: ["subjects"]
};
</script>
<style></style>