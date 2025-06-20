---
title: 广度优先搜索
date: 2022-06-03T15:00:00+08:00
draft: false
categories: [ "LeetCode", "算法学习" ]
tags: [ "LeetCode", "广度优先搜索", "BFS" ]
author: "swfoodt"
description: ""
---

### 简介

- 广度优先搜索（Breadth First Search）简称广搜或者 BFS.

- 广度优先搜索，感官上就像是水波的涟漪，从一个点开始，向外扩散，直到扩散到所有的点为止。下面这个例子 forked from [areaxe](https://github.com/Areaxe?tab=repositories)，可以很好的解释广度优先搜索的过程。

    <!-- <iframe height="900" width="100%" scrolling="no" title="bfs" src="https://codepen.io/swfoodt/embed/dyKyabG?default-tab=result" frameborder="no">
    </iframe> -->
  <!DOCTYPE html>
  <html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        #canvas {
            display: block;
            margin: 0 auto;
        }
    </style>
</head>

<body>

  <canvas id="canvas" width="600" height="600" />

  <script>
      let grid_width = 12; // 一个结点的宽度
      let canvas = document.getElementById('canvas')
      let ctx = canvas.getContext('2d');
      let originColor = 'coral'; // 搜索源点颜色
      let targetColor = 'cornflowerblue'; // 被搜索点颜色
      let visitedColor = 'rgb(255, 149, 142)'; // 已经遍历过的点的元素
      let dataWidth = 50; // 画板矩阵一行或者一列的点数
      let landSize = 5; // 源点个数
      let interval = 1000;
      let flag = 0;


      // 0表示待搜索的点，1表示源点，用2 表示已经访问过的点
      let grid = Array.from({
          length: dataWidth
      }, () => Array.from({
          length: dataWidth
      }, () => 0))

      window.onload = function init() {
          init_canvas()
      }


      function init_canvas() {
          // 初始化数据
          canvas.style.backgroundColor = targetColor;
          let originPositions = [];
          for (let i = 0; i < landSize; i++) {
              let randomx = parseInt(Math.random() * dataWidth)
              let randomy = parseInt(Math.random() * dataWidth)
              //  生成源点
              grid[randomx][randomy] = 1;
              originPositions.push([randomx, randomy]);
          }
          // initRender(grid,targetColor);  // 将所有点设置为被搜索的点
          ctx.beginPath();
          renderGrid(originPositions, originColor);
          ctx.save();
          wideSearch(originPositions)
      }

      function wideSearch(originPositions) {
          // let distance = -1;
          let queen = originPositions; // 待搜索队列
          let searchWidth = grid.length; // 搜索范围

          if (!queen.length || queen.length === searchWidth * searchWidth) {
              return -1;
          }

          let timer = setInterval(function () {
              let nextPosition = []; // 存储接下来将要渲染的结点位置 i j
              // 如果队列有结点
              if (queen.length) {
                  let pointLen = queen.length;
                  for (let i = 0; i < pointLen; i++) {
                      let position = queen.shift();
                      let x = position[0]; // 被搜索结点的横坐标
                      let y = position[1]; // 被搜索结点的纵坐标
                      // 向左搜索
                      if (x > 0 && grid[x - 1][y] === 0) {
                          queen.push([x - 1, y])
                          grid[x - 1][y] = 2;
                      }
                      //向右搜索
                      if (x < searchWidth - 1 && grid[x + 1][y] === 0) {
                          queen.push([x + 1, y])
                          grid[x + 1][y] = 2;
                      }
                      //向下搜索
                      if (y < searchWidth - 1 && grid[x][y + 1] === 0) {
                          queen.push([x, y + 1])
                          grid[x][y + 1] = 2;
                      }
                      //向上搜索
                      if (y > 0 && grid[x][y - 1] === 0) {
                          queen.push([x, y - 1])
                          grid[x][y - 1] = 2;
                      }
                  }
                  // 渲染
                  renderGrid(queen, visitedColor);
              } else {
                  clearInterval(timer)
              }
          }, interval)
      };

      // 渲染访问点
      function renderGrid(positionList, color) {
          let len = grid.length;
          for (let i = 0; i < positionList.length; i++) {
              let [x, y] = positionList[i];
              ctx.beginPath();
              ctx.moveTo(y * grid_width, x * grid_width);
              ctx.rect(y * grid_width, x * grid_width, grid_width, grid_width);
              ctx.fillStyle = color;
              ctx.fill()
              ctx.save()
          }
      }
  </script>

</body>

</html>

