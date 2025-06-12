const labels = JSON.parse(document.getElementById("labels").value);
const values = JSON.parse(document.getElementById("values").value);

const ctx = document.getElementById("eventChart").getContext("2d");

new Chart(ctx, {
  type: "bar",
  data: {
    labels: labels,
    datasets: [
      {
        label: "Event Attendance",
        data: values,
        backgroundColor: "lime",
        borderRadius: 5,
        barThickness: 30,
      },
    ],
  },
  options: {
    scales: {
      x: {
          offset: true, // <-- Add this line
        title: {
          display: true,
          text: "Events",
        },
        ticks: {
          autoSkip: false,
          maxRotation: 90,
          minRotation: 90,
        },
        grid: {
          display: false,
        },
        // 💡 This reduces spacing
        barPercentage: 0.4,
        categoryPercentage: 0.6,
      },
      y: {
        title: {
          display: true,
          text: "Event Attendance",
        },
        min: 0,
        max: 1000,
        ticks: {
          stepSize: 200,
        },
      },
    },
  },
});
