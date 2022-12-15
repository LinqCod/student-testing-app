<template>
  <div class="subjects">
    <h1>Предметы группы</h1>
    <subject-list :subjects="subjects"></subject-list>
  </div>
</template>
<script>
import SubjectList from "@/components/SubjectList.vue";
export default
{  name: "StudentSubjects",
  components: { SubjectList  },
  data: function() {
    return {
    };
  },
  computed: {
    subjects() {
      return this.$store.getters["subjects/allSubjects"];
    },
  },
  mounted() {
    this.$store.dispatch("subjects/getAllSubjects", { group_id: this.$store.getters["auth/currentStudent"].group.id });
  }
};
</script>