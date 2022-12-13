<template>
  <div id="app">
    <the-application-bar></the-application-bar>
    <router-view class="content" />
  </div>
</template>
<script>
import TheApplicationBar from "./components/UI/TheApplicationBar.vue";
export default {
  components: {    TheApplicationBar  }};
</script>
<style>
  body {
    margin: 0;
    color: rgba(255, 255, 255, 0.25);
    background-color: #0e1728;
    font-family: Arial, Helvetica, sans-serif;
  }
  .content {
    padding-top: 5.5vh;
    padding-left: 10%;
  }
</style>
<!--<script setup>-->
<!--  import { ref, computed } from 'vue'-->
<!--  //-->
<!--  // const headers = { "Content-Type": "application/json" };-->
<!--  //-->
<!--  // fetch("http://localhost:4000/api/v1/students/login", {-->
<!--  //   method: "POST",-->
<!--  //   headers,-->
<!--  //   mode: "no-cors",-->
<!--  //   body: JSON.stringify({-->
<!--  //     "email": "danny12123@yandex.ru",-->
<!--  //     "password": "1235gawfgbrtgw"-->
<!--  //   })-->
<!--  // })-->
<!--  //     .then(response => {-->
<!--  //       if (response.ok) {-->
<!--  //         return response.json();-->
<!--  //       } else throw Error(response.body);-->
<!--  //     })-->
<!--  //     .then(data => {-->
<!--  //       console.log("heelll")-->
<!--  //       console.log(data);-->
<!--  //     })-->
<!--  //     .catch(error => {-->
<!--  //       console.log(error);-->
<!--  //     });-->

<!--  // fetch("http://localhost:4000/api/v1/tasks/1", {-->
<!--  //   method: "GET",-->
<!--  //   headers,-->
<!--  //   mode: "no-cors"-->
<!--  // }).then(response => response.json())-->
<!--  //     .then(data => (console.log(data)))-->
<!--  //     .catch(err => console.log(err));-->

<!--  const questions = ref([-->
<!--    {-->
<!--      question: 'Q1',-->
<!--      answer: 0,-->
<!--      options: [-->
<!--          '1',-->
<!--          '2',-->
<!--          '3',-->
<!--          '4'-->
<!--      ],-->
<!--      selected: null,-->
<!--      index: null-->
<!--    },-->
<!--    {-->
<!--      question: 'Q2',-->
<!--      answer: 3,-->
<!--      options: [-->
<!--        '5',-->
<!--        '6',-->
<!--        '7',-->
<!--        '8'-->
<!--      ],-->
<!--      selected: null,-->
<!--      index: null-->
<!--    },-->
<!--    {-->
<!--      question: 'Q3',-->
<!--      answer: 2,-->
<!--      options: [-->
<!--        '9',-->
<!--        '10',-->
<!--        '11',-->
<!--        '12'-->
<!--      ],-->
<!--      selected: null,-->
<!--      index: null-->
<!--    },-->
<!--  ])-->

<!--  const quizCompleted = ref(false)-->
<!--  const currentQuestion = ref(0)-->
<!--  const score = computed(() => {-->
<!--    let value = 0-->
<!--    questions.value.map(q => {-->
<!--      if (q.selected != null && q.selected == q.answer) {-->
<!--        value++-->
<!--        console.log(value)-->
<!--      }-->
<!--    })-->
<!--    return value-->
<!--  })-->

<!--  const getCurrentQuestion = computed(() => {-->
<!--    let question = questions.value[currentQuestion.value]-->
<!--    question.index = currentQuestion.value-->
<!--    return question-->
<!--  })-->

<!--  const SetAnswer = event => {-->
<!--    questions.value[currentQuestion.value].selected = event.target.value-->
<!--    event.target.value = null-->
<!--  }-->

<!--  const NextQuestion = () => {-->
<!--    if (currentQuestion.value < questions.value.length - 1) {-->
<!--      currentQuestion.value++-->
<!--    } else {-->
<!--      quizCompleted.value = true-->
<!--    }-->
<!--  }-->
<!--</script>-->

<!--<template>-->
<!--  <main class="app">-->
<!--    <h1>Quiz app</h1>-->

<!--    <section class="quiz" v-if="!quizCompleted">-->
<!--      <div class="quiz-info">-->
<!--        <span class="question">{{ getCurrentQuestion.question }}</span>-->
<!--        <span class="score">Score {{ score }}/{{ questions.length }}</span>-->
<!--        <span class="answered">Answered {{ getCurrentQuestion.index }}/{{ questions.length }}</span>-->
<!--      </div>-->

<!--      <div class="options">-->
<!--        <label-->
<!--          v-for="(option, index) in getCurrentQuestion.options"-->
<!--          :key="index"-->
<!--          :class="`option ${-->
<!--            getCurrentQuestion.selected == index-->
<!--              ? index == getCurrentQuestion.answer-->
<!--                ? 'correct'-->
<!--                : 'wrong'-->
<!--              : ''-->
<!--          } ${-->
<!--            getCurrentQuestion.selected != null &&-->
<!--            index != getCurrentQuestion.selected-->
<!--              ? 'disabled'-->
<!--              : ''-->
<!--          }`">-->
<!--          <input-->
<!--              type="radio"-->
<!--              :name="getCurrentQuestion.index"-->
<!--              :value="index"-->
<!--              v-model="getCurrentQuestion.selected"-->
<!--              :disabled="getCurrentQuestion.selected"-->
<!--              @change="SetAnswer">-->
<!--          <span>{{ option }}</span>-->
<!--        </label>-->
<!--      </div>-->

<!--      <button-->
<!--        @click="NextQuestion"-->
<!--        :disabled="!getCurrentQuestion.selected">-->
<!--        {{-->
<!--          getCurrentQuestion.index == questions.length - 1-->
<!--            ? 'Finish'-->
<!--            : getCurrentQuestion.selected == null-->
<!--              ? 'Select an option'-->
<!--              : 'Next Question'-->
<!--        }}-->
<!--      </button>-->
<!--    </section>-->

<!--    <section v-else>-->
<!--      <h2>You have finished the quiz!</h2>-->
<!--      <p>Your score is {{ score }}/{{ questions.length }} </p>-->
<!--    </section>-->
<!--  </main>-->
<!--</template>-->

<!--<style>-->
<!--  * {-->
<!--    margin: 0;-->
<!--    padding: 0;-->
<!--    box-sizing: border-box;-->
<!--    font-family: 'Monospaced', sans-serif;-->
<!--  }-->

<!--  body {-->
<!--    background-color: #0a0909;-->
<!--    color: #FFF;-->
<!--  }-->

<!--  .app {-->
<!--    display: flex;-->
<!--    flex-direction: column;-->
<!--    align-items: center;-->
<!--    padding: 2rem;-->
<!--    min-height: 100vh;-->
<!--  }-->

<!--  h1 {-->
<!--    font-size: 2rem;-->
<!--    margin-bottom: 2rem;-->
<!--  }-->

<!--  .quiz {-->
<!--    background-color: #151313;-->
<!--    padding: 1rem;-->
<!--    width: 100%;-->
<!--    max-width: 640px;-->
<!--    border-radius: 0.5rem;-->
<!--  }-->

<!--  .quiz-info {-->
<!--    display: flex;-->
<!--    justify-content: space-between;-->
<!--    margin-bottom: 1rem;-->
<!--  }-->

<!--  .quiz-info .question {-->
<!--    color: rgba(255, 255, 255, 0.85);-->
<!--    font-size: 1.25rem;-->
<!--  }-->

<!--  .quiz-info .score,.answered {-->
<!--    color: rgba(255, 255, 255, 0.65);-->
<!--    font-size: 1.25rem;-->
<!--  }-->

<!--  .options {-->
<!--    margin-bottom: 1rem;-->
<!--  }-->

<!--  .option {-->
<!--    padding: 1rem;-->
<!--    display: block;-->
<!--    background-color: #0a0909;-->
<!--    margin-bottom: 0.5rem;-->
<!--    border-radius: 0.5rem;-->
<!--    cursor: pointer;-->
<!--  }-->

<!--  .option:hover {-->
<!--    background-color: #ffffff;-->
<!--    color: black;-->
<!--  }-->

<!--  .option.correct {-->
<!--    background-color: chartreuse;-->
<!--  }-->

<!--  .option.wrong {-->
<!--    background-color: red;-->
<!--  }-->

<!--  .option.last-of-type {-->
<!--    margin-bottom: 0;-->
<!--  }-->

<!--  .option.disabled {-->
<!--    opacity: 0.5;-->
<!--  }-->

<!--  .option input {-->
<!--    display: none;-->
<!--  }-->

<!--  button {-->
<!--    appearance: none;-->
<!--    outline: none;-->
<!--    border: none;-->
<!--    cursor: pointer;-->

<!--    padding: 0.5rem 1rem;-->
<!--    background-color: #41ff41;-->
<!--    color: #000000;-->
<!--    font-weight: 700;-->
<!--    text-transform: uppercase;-->
<!--    font-size: 1.25rem;-->
<!--    border-radius: 0.5rem;-->
<!--  }-->

<!--  button:disabled {-->
<!--    opacity: 0.5;-->
<!--  }-->

<!--  h2 {-->
<!--    font-size: 2rem;-->
<!--    margin-bottom: 2rem;-->
<!--    text-align: center;-->
<!--  }-->

<!--  p {-->
<!--    color: #afafaf;-->
<!--    font-size: 1.25rem;-->
<!--    text-align: center;-->
<!--  }-->
<!--</style>-->
