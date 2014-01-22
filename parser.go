package woothee

import (
  "strings"
)

type Parser struct {
  AgentDataSet DataSet
}

func NewParser() (*Parser) {
  return &Parser { DefaultDataSet }
}

func (p *Parser) LookupDataset(label string) (*Result, error) {
  result := p.AgentDataSet[label]
  if result == nil {
    return nil, ErrNoDataSet
  }

  return result, nil
}

func (p *Parser) TryCrawler(agent string) (result *Result, err error) {
  result, err = p.ChallengeGoogle(agent)
  if err == nil {
    return result, nil
  }

  result, err = p.ChallengeCrawlers(agent)
  if err == nil {
    return result, nil
  }

  return nil, ErrNoMatch
}


func (p *Parser) ChallengeGoogle(agent string) (*Result, error) {
  if ! strings.Contains(agent, "Google") {
    return nil, ErrNoMatch
  }

  if strings.Contains(agent, "compatible; Googlebot") {
    if strings.Contains(agent, "compatible; Googlebot-Mobile") {
      return p.LookupDataset("GoogleBotMobile")
    } else {
      return p.LookupDataset("GoogleBot")
    }
  }

  if strings.Contains(agent, "Googlebot-Image/") {
    return p.LookupDataset("GoogleBot")
  }

  return nil, ErrNoMatch
}

func (p *Parser) ChallengeCrawlers(agent string) (*Result, error) {
  return nil, ErrNoMatch
}
