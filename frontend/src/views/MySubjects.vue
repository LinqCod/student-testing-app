<template>
  <div class="subjects">
    <h1>Предметы группы</h1>
    <subject-list :subjects="subjects"></subject-list>
  </div>
</template>
<script>
import SubjectList from "@/components/SubjectList.vue";
export default
{  name: "MySubjects",
  components: { SubjectList  },
  data: function() {
    return {
      subjects: [
        {
          id: 1,
          title: "Математика"
        },
        {
          id: 2,
          title: "Программирование"
        }
      ]
    };
  }};
</script>