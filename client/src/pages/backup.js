
import React, { useEffect, useState } from 'react'
import axios from 'axios';
import {
  Accordion, AccordionSummary,
  AccordionDetails, Typography,
  Radio, RadioGroup, FormControlLabel,
  FormControl, FormLabel, Button,
  Container, Grid
} from '@material-ui/core'
import ExpandMoreIcon from '@material-ui/icons/ExpandMore'

export const UserLandingPage = () => {
  const [expandedPanel, setExpandedPanel] = useState(false);
  const [polls, setPolls] = useState([]);
  const [votedPollSummary, setvotedPollSummary] = useState([]);
  const [selectedPollId, setSelectedPollId] = useState(-1);
  const [selectedPollOption, setselectedPollOption] = useState("");
  const [errorMessage, setErrorMessage] = useState('');
  const [pollTopic, setPollTopicValue] = useState('');
  const [pollOptions, setpollOptions] = useState([]);

  const userId = localStorage.getItem('userId');
  const token = localStorage.getItem('token');

  useEffect(() => {
    (async () => {
      const response = await axios({
        method: 'get',
        url: '/poll/' + userId,
        headers: {
          'Content-Type': 'multipart/form-data',
          'Authorization': `Bearer ${token}`
        }
      });
      // console.log(response.data.polls)
      console.log("response.data.polls type is Array : ", Array.isArray(response.data.polls))
      console.log("response.data.polls type is Array : ", Array.isArray(response.data.votedPollSummary))
      const newPolls = [...response.data.polls];
      setPolls(newPolls)
      const votedPolls = [...response.data.votedPollSummary];

      const transformed = Object.entries(votedPolls.reduce((acc, { option, count, pollid }) => {
        acc[pollid] = (acc[pollid] || []);
        acc[pollid].push({ option, count });
        return acc;
      }, {})).map(([key, value]) => ({ pollid: key, votes: value }));

      console.log("tarnsformed data " + transformed);
      // setvotedPollSummary(votedPolls)
      setvotedPollSummary(transformed)
    })();
  }, [])
  const handleVoteChange = e => {
    setselectedPollOption(e.target.value)
    alert("selectedPollId : " + selectedPollId + " selectedPollOption : " + e.target.value)
  }

  const handleVote = async () => {
    var bodyFormData = new FormData();

    bodyFormData.append('voterId', userId);
    bodyFormData.append('pollId', selectedPollId);
    bodyFormData.append('option', selectedPollOption);
    const response = await axios({
      method: 'post',
      url: '/poll/vote',
      data: bodyFormData,
      headers: {
        'Content-Type': 'multipart/form-data',
        'Authorization': `Bearer ${token}`
      }
    });
    const { votingId } = response.data;
    alert("Vote registerd with votingId : " + votingId);

  }

  //Create Poll 
  const onSubmitPollClicked = async () => {
    var bodyFormData = new FormData();
    bodyFormData.append('ownerId', userId);
    bodyFormData.append('topic', pollTopic);
    bodyFormData.append('options', JSON.stringify(pollOptions));

    const response = await axios({
      method: 'post',
      url: '/poll',
      data: bodyFormData,
      headers: {
        'Content-Type': 'multipart/form-data',
        'Authorization': `Bearer ${token}`
      }
    });
    const { pollId } = response.data;
    alert("Poll create with pollId : " + pollId);

  }

  // Create poll set options
  const handleSetOptions = (content) => {
    console.log("content : ", content);
    const options = content.split("|");
    setpollOptions(options)
  }
  const handleAccordionChange = (pollId) => (event, isExpanded) => {
    // console.log({ event, isExpanded });
    console.log("pollId : " + pollId)
    setExpandedPanel(isExpanded ? pollId : false);
    setSelectedPollId(pollId);
  };

  return (
    // <Accordion expanded={expandedPanel === 'panel1'} onChange={handleAccordionChange('panel1')}>
    // <div>
    <Container>
      <Grid container>
        <h1>New Polls</h1>
      </Grid>
      <Grid container spacing={3}>

        {polls.length > 0 ? polls.map((poll) => (
          <Grid item xs={12} md={6} lg={4} key={poll.id}>
            <Accordion expanded={expandedPanel === poll.id} onChange={handleAccordionChange(poll.id)}>
              <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel1a-content"
                id="panel1a-header"
              >
                <Typography><h2>{poll.topic}</h2></Typography>
              </AccordionSummary>
              <AccordionDetails>
                <FormControl component="fieldset">
                  <FormLabel component="legend">Choose One</FormLabel>
                  <RadioGroup row aria-label="choose" name="row-radio-buttons-group"
                    onChange={handleVoteChange} value={selectedPollOption}>
                    {poll.options.map((option) => (
                      <FormControlLabel value={option} control={<Radio />} label={option} />
                    ))}
                  </RadioGroup>
                  <Button variant="contained" size="small" onClick={handleVote}>
                    Vote
                  </Button>
                </FormControl>
              </AccordionDetails>
            </Accordion>
          </Grid>
        )) : null}
      </Grid>
      <Grid container>
        <h1>You have Voted Already ..</h1>
      </Grid>
      <Grid container spacing={3}>
        {votedPollSummary.length > 0 ? votedPollSummary.map((summary) => (
          <Grid item xs={12} md={6} lg={4} key={summary.pollid}>
            <h3>{summary.pollid}</h3>
            {summary.votes.length > 0 ? summary.votes.map((vote) => (
              <p>{vote.option}   :   {vote.count} </p>
            )) : null}
          </Grid>
        )) : null}
      </Grid>
      <Grid container>
        <h1>Create Poll</h1>
      </Grid>
      <Grid container>
        <Accordion>
          <AccordionSummary
            expandIcon={<ExpandMoreIcon />}
            aria-controls="panel1a-content"
            id="panel1a-header"
          >
            <Typography><h2>Create Poll</h2></Typography>
          </AccordionSummary>
          <AccordionDetails>
            <div className="content-container">
              <h1>Create Poll</h1>
              {errorMessage && <div className="fail">{errorMessage}</div>}
              <input
                value={pollTopic}
                onChange={e => setPollTopicValue(e.target.value)}
                placeholder="Provide a name for the poll you want to create" />
              <input
                value={pollOptions}
                onChange={e => handleSetOptions(e.target.value)}
                placeholder="pipe '|' separated options" />
              <hr />
              <button
                disabled={!pollTopic || !pollOptions}
                onClick={onSubmitPollClicked}>Submit</button>
            </div>
          </AccordionDetails>
        </Accordion>

        {/* <FormControl component="fieldset">
          <FormLabel component="legend">Create a Poll</FormLabel>
              
          <Button variant="contained" size="small" onClick={handleCreatePoll}>
            Create
          </Button>
        </FormControl> */}
      </Grid>
    </Container>

  );

}