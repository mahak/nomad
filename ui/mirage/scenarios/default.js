export default function(server) {
  server.createList('job', 10);
  server.createList('node', 10);
  server.createList('agent', 3);
  server.createList('allocation', 20);
}
