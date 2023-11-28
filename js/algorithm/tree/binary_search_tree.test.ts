import { BinarySearchTree } from "./binary_search_tree"; // adjust the import path as needed

describe("BinarySearchTree", () => {
  let bst: BinarySearchTree;

  beforeEach(() => {
    bst = new BinarySearchTree();
  });

  test("insert and search", () => {
    bst.insert(5);
    bst.insert(3);
    bst.insert(7);
    expect(bst.searchNode(bst.root, 5)).toBe(true);
    expect(bst.searchNode(bst.root, 3)).toBe(true);
    expect(bst.searchNode(bst.root, 7)).toBe(true);
    expect(bst.searchNode(bst.root, 10)).toBe(false);
  });

  test("delete", () => {
    bst.insert(5);
    bst.insert(3);
    bst.insert(7);
    bst.deleteNode(bst.root, 3);
    expect(bst.searchNode(bst.root, 3)).toBe(false);
  });
});
