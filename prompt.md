Analyze the documentation for the adk project and please come up with a list of features that adk delivers. Focus on page titles and section headers to come up with that list.

Once that is done, I'd like to generate a matrix of features enabled by runtime.

Currently there are two runtimes: python (adk-python) which is the primary, and java (adk-java). But soon go (adk-go) and typescript (adk-js) will be added.

A good metric for figuring out if a feature is supported per runtime is if the sample code associated with a given feature in adk-docs includes samples in a given language. 

For each matrix item showing support, please link to the code sample where you use to say that this feature is supported in the language. Use https://google.github.io/adk-docs/ as the base link. Also mkae sure to include path and page so the links will work. 

Output the content in markdown. 

Please double check the matrix's markdown content for validity before outputting. 

Please say "Planned" instead of "No" when a feature is not supported, unless there is text in the documentation that makes it clear that a particular feature will not be supported in a given language, in which case, then write "no". If you mark something as know please include a link to the text that makes you make that determination. Nearest header is close enough. 

Write the output as if it was going to be a page available in adk-docs. Match adk-docs style and tone. 

For now omit Typescript and Go from matrix. 

Always preface the content with ********** so that I can programmatically remove your commentary from the content. 

Pattern for building urls from markdown paths should be to replaced the .md extension with a "/". 

Use the following markdown as an example of a good previous run of this content and should really only be updated if content in adk-docs has changed: