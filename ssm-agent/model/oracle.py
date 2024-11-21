#!/usr/bin/env python3

from sys import argv, stdin, exit
import re

def splitter(line):
  splitted = line.split(':')
  return (splitted[0], splitted[1].strip())

lines = list(map(splitter, stdin.readlines()))
if not lines:
  exit(0)

def subToken(token, line):
  (num, goal) = line
  if isinstance(token, str):
    return num if token in goal else None
  else:
    return num if token.search(goal) is not None else None

def matchAgainstList(priorityList, lines):
  for token in priorityList:
    try:
      return next(filter(bool, map(lambda line: subToken(token, line), lines)))
    except StopIteration:
      pass

def filterNone(list):
  return [x for x in list if x is not None]

match = None
if argv[1] in ['loop_induction_client']:
  match = matchAgainstList([
    'Client_S7',
  ], lines)
elif argv[1] in ['loop_induction_agent']:
  match = matchAgainstList([
    'St_Agent_10',
  ], lines)
elif argv[1] in ['x_is_secret']:
  match = matchAgainstList([
    'St_Agent_5',
    '!KU( ~x )',
  ], lines)
elif argv[1] in ['y_is_secret']:
  match = matchAgainstList([
    'Client_S5',
    '!KU( ~y )',
  ], lines)
elif argv[1] in ['AgentSendEncryptedSessionKey_is_unique']:
  match = matchAgainstList([
    'St_Agent_8',
  ], lines)
elif argv[1] in ['SharedSecret_is_secret','injectiveagreement_agent','injectiveagreement_client']:
  match = matchAgainstList(filterNone([
    '!KU( ~x )',
    '!KU( ~y )',
    re.compile(r'!KU\( kdf1\(\'g\'\^\(~?[xy]\*~?[xy]\)\) \)'),
    re.compile(r'!KU\( \'g\'\^\(~[xy]\*'),
    '!KU( \'g\'^inv(~x) )',
    '!KU( \'g\'^inv(~y) )',
    re.compile(r'!KU\( \'g\'\^\(.*inv\(.*\)\)'),
    '!KU( ~ltk )',
    '!KU( ~x.',
    '!KU( ~y.',
    '(∀ #j. (K( <kdf1(z), kdf2(z)> ) @ #j) ⇒ ⊥)', # secrecy lemma
    'St_Agent_9',
    'Client_S6',
    '<\'SignResponse\'',
    '\'VerifyResponse\'',
    '<\'VerifyRequest\'',
    '<\'SignRequest\'',
    '!KU( senc(<\'HandshakeCompletePayload\'' if argv[1] in ['SharedSecret_is_secret'] else None,
    '!KU( sign(<Y,',
    '!KU( sign(<X,',
    '!KU( ~SignKey',
    '!KU( sign(<\'g\'^(~x*',
    '!KU( sign(<\'g\'^x,' if argv[1] in ['injectiveagreement_client'] else None,
    'splitEqs(',
  ]), lines)

if match is not None:
  print(match)
