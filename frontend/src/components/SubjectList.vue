<template>
  <div class="subject-list">
    <single-subject v-for="subject in subjects" :key="subject.id" :subject="subject" @click.prevent="subjectClick(subject)">
    </single-subject>
  </div>
</template>
<script>
import SingleSubject from "./SingleSubject.vue";
export default {
  components: { SingleSubject },
  props: ["subjects"],
  methods: {
    subjectClick(subject) {
      this.$router.push({ name: "TaskSections", params: { subject_id: subject.id } });
    }
  }
};
</script>
<style></style>