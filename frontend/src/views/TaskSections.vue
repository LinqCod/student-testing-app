<template>
  <div class="sections">
    <h1>Разделы предмета</h1>
    <section-list :sections="sections"></section-list>
  </div>
</template>
<script>
import SectionList from "@/components/SectionList.vue";
export default
{  name: "TaskSections",
  components: { SectionList  },
  data: function() {
    return {
    };
  },
  computed: {
    sections() {
      return this.$store.getters["sections/allSections"];
    }
  },
  props: [
    "subject_id"
  ],
  mounted() {
    this.$store.dispatch("sections/getAllSections", { subject_id: this.subject_id });
  }
};
</script>