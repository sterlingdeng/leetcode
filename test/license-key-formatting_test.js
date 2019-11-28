const assert = require("assert");
const {
  licenseKeyFormatting
} = require("../license_key_formatting/license-key-formatting");

describe("license key formatting", () => {
  const testCases = [
    {
      name: "1",
      in: "5F3Z-2e-9-w",
      k: 4,
      expect: "5F3Z-2E9W"
    },
    {
      name: "2",
      in: "2-5g-3-J",
      k: 2,
      expect: "2-5G-3J"
    },
    {
      name: "3",
      in: "2-5g-3-J",
      k: 1,
      expect: "2-5-G-3-J"
    },
    {
      name: "4",
      in: "2-5g-3-J",
      k: 7,
      expect: "25G3J"
    }
  ];
  testCases.forEach(testCase => {
    it(testCase.name, () => {
      const got = licenseKeyFormatting(testCase.in, testCase.k);
      assert.deepEqual(
        got,
        testCase.expect,
        `licenseKeyFormatting() = ${got}, but expected ${testCase.expect}`
      );
    });
  });
});
