import Highlight, { defaultProps } from "prism-react-renderer";
import styled from "styled-components";
import theme from "prism-react-renderer/themes/vsLight";

interface IProps {
  code?: string;
  language: string;
}

const Pre = styled.pre`
  text-align: left;
  //   margin: 1em 0;
  //   padding: 0.5em;
  overflow: scroll;

  & .token-line .token {
    line-height: 1.3em;
    // height: 1.3em;
    white-space: pre-wrap;
    word-break: break-word;
  }
`;

export default function Code({ code, language }: IProps) {
  return (
    <Highlight
      {...defaultProps}
      theme={theme}
      code={code ? code : ""}
      language="json"
    >
      {({ className, style, tokens, getLineProps, getTokenProps }) => (
        <Pre className={className} style={style}>
          {tokens.map((line, i) => (
            <div {...getLineProps({ line, key: i })} className="text-sm">
              {line.map((token, key) => (
                <span {...getTokenProps({ token, key })} />
              ))}
            </div>
          ))}
        </Pre>
      )}
    </Highlight>
  );
}
