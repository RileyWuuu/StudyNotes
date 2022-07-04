# Camel2Snake converting

## Question:
* read json file with camel case key field and write out converted key field name as snake_case
* using stand library only. don’t use any 3rd party library/framework
* allowed language: golang

## Solution:
1. Unmarshal json data
2. Cut string with regexp.Mustcompile
3. run data with for range to convert the camel case key into snake case
4. In order to convert all key fields in nested structure, use switch to get each type 
5. By using recursion, we're allowed to run through the whole nested-structure data and convert all keyfields.
