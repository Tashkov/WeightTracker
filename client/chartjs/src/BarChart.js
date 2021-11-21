import React from "react";
import { Line, Bar } from "react-chartjs-2";

let jsonData = require('./data.json');
let filData = require('./fil_data.json')
console.log(jsonData.length);
// jsonData = jsonData.slice(0, 100);

let reducedData = [];
jsonData.reduce((res, value) => {
    if (!res[value.calendar_last_scraped]) {
        res[value.calendar_last_scraped] = {
            calendar_last_scraped: value.calendar_last_scraped,
            accommodates: 0
        };
        reducedData.push(res[value.calendar_last_scraped]);
    }
    res[value.calendar_last_scraped].accommodates += value.accommodates;
    return res;
}, {});
jsonData = reducedData;
jsonData = jsonData.sort((a, b) => (a.calendar_last_scraped > b.calendar_last_scraped) ? 1 : -1);
jsonData = filData;
const xAxis = jsonData.map(item => item.calendar_last_scraped);
const yAxis = jsonData.map(item => item.accommodates);

const lineData = {
    labels: xAxis,
    datasets: [{
        label: '# of accomodations',
        data: yAxis,
        backgroundColor: [
            'rgba(255, 99, 132, 0.2)'
        ],
        borderColor: [
            'rgba(255, 99, 132, 1)',
        ],
        fill: false,
        borderWidth: 1,
        tension: 0,
        elements: {
            point: {
                radius: 0
            }
        }
    }]
};

const BarChart = ({ data, options }) => {
    return (
        <div>
            <Line
                data={lineData}
                height={400}
                width={600}
            />
            <canvas id="myChart" />
        </div>
    );
};
export default BarChart;