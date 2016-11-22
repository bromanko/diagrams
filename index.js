import './scss/main.scss';

import Viz from 'viz.js';

const src = 'digraph {One -> Two -> Three;}';

const graphSvg = Viz(src, {
  format: 'svg',
  engine: 'dot',
});

const container = document.getElementById('graph-holder');
container.innerHTML = graphSvg;

const svgEl = container.getElementsByTagName('svg')[0];
svgEl.id = 'graph';
svgEl.style.width = '100%';
svgEl.style.height = '100%';
