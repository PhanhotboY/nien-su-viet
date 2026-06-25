import { toSnakeCase } from '../util';

const getRoutingKey = function (eventName: string) {
  return toSnakeCase(eventName);
};

const getDeadRoutingKey = function (eventName: string) {
  return `${toSnakeCase(eventName)}_dead`;
};

const getQueueName = function (eventName: string) {
  return `${toSnakeCase(eventName)}_queue`;
};

const getDLXName = function (eventName: string) {
  return `${toSnakeCase(eventName)}_dlx`;
};

const getDLQName = function (eventName: string) {
  return `${toSnakeCase(eventName)}_dlq`;
};

export {
  getRoutingKey,
  getDeadRoutingKey,
  getQueueName,
  getDLXName,
  getDLQName,
};
