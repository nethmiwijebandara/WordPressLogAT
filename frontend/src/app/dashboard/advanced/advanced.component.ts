import { Component, OnInit } from '@angular/core';
import { Chart } from 'chart.js';

@Component({
  selector: 'app-advanced',
  templateUrl: './advanced.component.html',
  styleUrls: ['./advanced.component.scss']
})
export class AdvancedComponent implements OnInit {

  barChart: any;
  pieChart: any;
  lineChart: any;

  constructor() { }

  ngOnInit(): void {
    this.showBar();
    this.showPie();
    this.showLine();
  }

  showBar(): void {
    this.barChart = new Chart('barChart', {
      type: 'bar',
      data: {
        labels: ['Health', 'IT', 'Retail', 'Financial', 'Logistics', 'Automobile', 'Food', 'Transport'],
        datasets: [{
          label: '',
          data: [12, 19, 3, 5, 2, 3, 10, 15],
          backgroundColor: '#27AE60',
        }
        ]
      },
      options: {
        legend: {
          display: false
        }
      }
    });
  }

  showPie(): void {
    this.pieChart = new Chart('pieChart', {
      type: 'doughnut',
      data: {
        labels: ['Designer', 'Developer', 'Executive', 'Manager'],
        datasets: [{
          label: '',
          data: [30, 10, 8, 20],
          backgroundColor: ['#F1C40F', '#2980B9', '#8E44AD', '#E74C3C']
        }]
      }
    });
  }

  showLine(): void {
    this.lineChart = new Chart('lineChart', {
      type: 'line',
      data: {
        labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
        datasets: [{
          lebel: '',
          data: [10, 15, 5, 8, 20, 18, 25, 6, 7, 10, 22, 3],
          backgroundColor: '#34495E',
          fill: false,
          lineTension: 0,
          borderColor: '#34495E',
          borderCapStyle: 'butt',
          borderDashOffset: 0.0,
          borderJoinStyle: 'miter',
          pointBorderColor: '#34495E',
          pointBackgroundColor: '#fff',
          pointBorderWidth: 1,
          pointHoverRadius: 7,
          pointHoverBackgroundColor: '#ffffff',
          pointHoverBorderColor: '#34495E',
          pointHoverBorderWidth: 2,
          pointRadius: 2,
          pointHitRadius: 5,
          spanGaps: false
        }]
      },
      options: {
        legend: {
          display: false
        }
      }
    });
  }

}
