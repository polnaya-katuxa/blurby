<template>
  <BaseSection>

    <div class="card-tr mb-2">
      <div class="row">
        <div style="align-items: center; display: flex" class="col-3">
          <img src="../assets/logo512.png" width="306" height="155" alt="image"
               class="center mx-4 px-1"/>
        </div>
        <div class="col-9 text-center">
          <p class="mt-3 nav-link" style="color: #ffffff; font-weight: bolder; font-size: 2em">
            Hello, {{ curLogin }} :) </p>
          <p class="nav-link" style="color: #ffffff; font-weight: bold; font-size: 1em">
            Create, customize, and conquer with <br>
            our targeted messaging service. </p>
        </div>
      </div>
    </div>

    <div class="wrapper">
      <div class="card text-center mb-2 box1">
        <div class="card-body">
          <h1 class="card-title gradient-text2" style="font-size: 5em">{{ clientStats.num }}</h1>
          <p class="card-text">Number of clients
          </p>
        </div>
      </div>

      <div class="card text-center mb-2 box2">
        <div class="card-body">
          <h1 class="card-title gradient-text3" style="font-size: 5em">{{ clientStats.avgAge }}</h1>
          <p class="card-text">Average client age
          </p>
        </div>
      </div>

      <div class="card text-center mb-2 ms-2 box3">
        <div class="card-body">
          <Line style="height: 100% !important;"
            id="my-chart-id"
            :options="chartOptions"
            :data="chartData"
          />
        </div>
      </div>
    </div>

    <div class="card-tr mb-2 mt-1">
      <p class="nav-link mt-2" style="color: #ffffff; font-weight: bold; font-size: 1em">
        Please, make sure your clients agreed to share personal information!</p>
    </div>

  </BaseSection>
</template>

<script lang="ts">
// eslint-disable-next-line max-classes-per-file
import { defineComponent } from 'vue';

import { mapActions, mapGetters } from 'vuex';
import BaseSection from '@/components/BaseSection.vue';
import { Line } from 'vue-chartjs';
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Filler,
} from 'chart.js';

ChartJS.register(
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Filler,
);

export default defineComponent({
  name: 'HomeView',
  components: { BaseSection, Line }, // , StatCard },
  mounted() {
    this.getStats();
  },
  methods: {
    ...mapActions(['getStats', 'userInfo']),
  },
  computed: {
    ...mapGetters(['clientStats', 'adStats', 'curLogin', 'chartAdNums', 'chartAdDates']),
    chartData() {
      return {
        labels: this.chartAdDates,
        datasets: [
          {
            label: 'Ad sends',
            type: 'line',
            data: this.chartAdNums,
            fill: true,
            borderColor: 'rgb(236,158,218)',
            backgroundColor: 'rgba(253,187,235,0.5)',
            borderWidth: 2,
            tension: 0.2,
          },
        ],
      };
    },
    chartOptions() {
      return {
        responsive: true,
        maintainAspectRatio: false,
        color: 'black',
        plugins: {
          legend: {
            labels: {
              font: {
                size: 14,
                family: 'Ubuntu',
              },
            },
          },
        },
      };
    },
  },
});
</script>
