import _ from 'lodash';

function omit<T = Object>(obj: T, fields?: string[]): Partial<T> {
  if (!fields || !fields.length) return obj;

  return (Object.keys(obj as Object) as Array<keyof typeof obj>).reduce<
    Partial<T>
  >((object, field) => {
    if (!fields.includes(field as string)) object[field] = obj[field];

    return object;
  }, {});
}

const isNullish = (val: any) => (val ?? null) === null;
const isEmptyObj = (obj: Object) => !Object.keys(obj).length;
const getSkipNumber = (limit: number, page: number) => limit * (page - 1);
const isObj = (obj: any) => obj instanceof Object && !Array.isArray(obj);
const capitalizeFirstLetter = (str: string) =>
  str.charAt(0).toUpperCase() + str.slice(1);

function removeNullishElements(arr: Array<any>) {
  const final: typeof arr = [];

  arr.forEach((ele) => {
    if (!isNullish(ele) && ele !== '') {
      const result = removeNestedNullish(ele);
      if (result instanceof Object && isEmptyObj(result)) return;

      final[final.length] = result;
    }
  });

  return final.filter((ele) => !isNullish(ele) && ele);
}

function removeNullishAttributes(obj: Record<string, any>) {
  const final: typeof obj = {};

  (Object.keys(obj) as Array<keyof typeof obj>).forEach((key) => {
    if (!isNullish(obj[key]) && obj[key] !== '') {
      const result = removeNestedNullish(obj[key]);

      if (result instanceof Object && isEmptyObj(result)) return;

      final[key] = result;
    }
  });

  return final;
}

function removeNestedNullish<T>(any: any): T {
  if (any instanceof Array)
    return removeNullishElements(any as Array<any>) as T;
  if (any instanceof Object) return removeNullishAttributes(any as Object) as T;

  return any;
}

function removeNestedUndefined<T>(any: any): T {
  if (any instanceof Array)
    return removeUndefinedElements(any as Array<any>) as T;
  if (any instanceof Object)
    return removeUndefinedAttributes(any as Object) as T;

  return any;
}

function removeUndefinedAttributes(obj: Record<string, any>) {
  const final: typeof obj = {};

  (Object.keys(obj) as Array<keyof typeof obj>).forEach((key) => {
    if (obj[key] !== undefined) {
      const result = removeNestedUndefined(obj[key]);

      if (result instanceof Object && isEmptyObj(result)) return;

      final[key] = result;
    }
  });

  return final;
}

function removeUndefinedElements(arr: Array<any>) {
  const final: typeof arr = [];

  arr.forEach((ele) => {
    if (ele !== undefined) {
      const result = removeNestedUndefined(ele);
      if (result instanceof Object && isEmptyObj(result)) return;

      final[final.length] = result;
    }
  });

  return final.filter((ele) => ele !== undefined);
}

export {
  omit,
  isNullish,
  isEmptyObj,
  getSkipNumber,
  isObj,
  capitalizeFirstLetter,
  removeNestedNullish,
  removeNestedUndefined,
};
