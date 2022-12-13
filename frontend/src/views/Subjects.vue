<template>
  <div class="subjects">
    <h1>Subjects</h1>
    <subjects-list :subjects="subjects"></subjects-list>
  </div>
</template>
<script>
import SubjectsList from "@/components/SubjectsList.vue";
export default
{  name: "Subjects",
  components: { SubjectsList  },
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