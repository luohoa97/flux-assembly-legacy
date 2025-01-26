// Generated from d:/Programs/Programmed/fluxlang/antlr4/FLUXASSEMBLY/src/lua_grammar_antlr4.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class lua_grammar_antlr4Parser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, NUMBER=30, STRING=31, 
		LETTER=32, DIGIT=33, WS=34;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_assignment = 2, RULE_expression = 3, 
		RULE_literal = 4, RULE_function_call = 5, RULE_table_insert = 6, RULE_function_declaration = 7, 
		RULE_block = 8, RULE_if_statement = 9, RULE_for_statement = 10, RULE_while_statement = 11, 
		RULE_table = 12, RULE_field = 13, RULE_table_access = 14, RULE_multi_line_comment = 15, 
		RULE_identifier = 16;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "assignment", "expression", "literal", "function_call", 
			"table_insert", "function_declaration", "block", "if_statement", "for_statement", 
			"while_statement", "table", "field", "table_access", "multi_line_comment", 
			"identifier"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'local'", "'='", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'true'", 
			"'false'", "'nil'", "','", "'table.insert'", "'function'", "'end'", "'if'", 
			"'then'", "'elseif'", "'else'", "'for'", "'do'", "'while'", "'{'", "'}'", 
			"'['", "']'", "'--[['", "'--]]'", "']]'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, "NUMBER", "STRING", "LETTER", "DIGIT", 
			"WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "lua_grammar_antlr4.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public lua_grammar_antlr4Parser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<Multi_line_commentContext> multi_line_comment() {
			return getRuleContexts(Multi_line_commentContext.class);
		}
		public Multi_line_commentContext multi_line_comment(int i) {
			return getRuleContext(Multi_line_commentContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(38);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 7664135690L) != 0)) {
				{
				setState(36);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case T__0:
				case T__2:
				case T__8:
				case T__9:
				case T__10:
				case T__12:
				case T__13:
				case T__15:
				case T__19:
				case T__21:
				case T__22:
				case NUMBER:
				case STRING:
				case LETTER:
					{
					setState(34);
					statement();
					}
					break;
				case T__26:
					{
					setState(35);
					multi_line_comment();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(40);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public For_statementContext for_statement() {
			return getRuleContext(For_statementContext.class,0);
		}
		public While_statementContext while_statement() {
			return getRuleContext(While_statementContext.class,0);
		}
		public Function_declarationContext function_declaration() {
			return getRuleContext(Function_declarationContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(47);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(41);
				assignment();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(42);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(43);
				if_statement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(44);
				for_statement();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(45);
				while_statement();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(46);
				function_declaration();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_assignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(50);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__0) {
				{
				setState(49);
				match(T__0);
				}
			}

			setState(52);
			identifier();
			setState(53);
			match(T__1);
			setState(54);
			expression(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public Token arith_operator;
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public Function_callContext function_call() {
			return getRuleContext(Function_callContext.class,0);
		}
		public TableContext table() {
			return getRuleContext(TableContext.class,0);
		}
		public Table_accessContext table_access() {
			return getRuleContext(Table_accessContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 6;
		enterRecursionRule(_localctx, 6, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				{
				setState(57);
				literal();
				}
				break;
			case 2:
				{
				setState(58);
				identifier();
				}
				break;
			case 3:
				{
				setState(59);
				function_call();
				}
				break;
			case 4:
				{
				setState(60);
				table();
				}
				break;
			case 5:
				{
				setState(61);
				table_access();
				}
				break;
			case 6:
				{
				setState(62);
				match(T__2);
				setState(63);
				expression(0);
				setState(64);
				match(T__3);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(76);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(74);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(68);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(69);
						((ExpressionContext)_localctx).arith_operator = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__4 || _la==T__5) ) {
							((ExpressionContext)_localctx).arith_operator = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(70);
						expression(3);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(71);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(72);
						((ExpressionContext)_localctx).arith_operator = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__6 || _la==T__7) ) {
							((ExpressionContext)_localctx).arith_operator = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(73);
						expression(2);
						}
						break;
					}
					} 
				}
				setState(78);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ParserRuleContext {
		public TerminalNode NUMBER() { return getToken(lua_grammar_antlr4Parser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(lua_grammar_antlr4Parser.STRING, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(79);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 3221229056L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Function_callContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public Table_insertContext table_insert() {
			return getRuleContext(Table_insertContext.class,0);
		}
		public Function_callContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_call; }
	}

	public final Function_callContext function_call() throws RecognitionException {
		Function_callContext _localctx = new Function_callContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_function_call);
		int _la;
		try {
			setState(96);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LETTER:
				enterOuterAlt(_localctx, 1);
				{
				setState(81);
				identifier();
				setState(82);
				match(T__2);
				setState(91);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 7524593160L) != 0)) {
					{
					setState(83);
					expression(0);
					setState(88);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__11) {
						{
						{
						setState(84);
						match(T__11);
						setState(85);
						expression(0);
						}
						}
						setState(90);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(93);
				match(T__3);
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 2);
				{
				setState(95);
				table_insert();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Table_insertContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Table_insertContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_table_insert; }
	}

	public final Table_insertContext table_insert() throws RecognitionException {
		Table_insertContext _localctx = new Table_insertContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_table_insert);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(98);
			match(T__12);
			setState(99);
			match(T__2);
			setState(100);
			identifier();
			setState(101);
			match(T__11);
			setState(102);
			expression(0);
			setState(103);
			match(T__3);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Function_declarationContext extends ParserRuleContext {
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public Function_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_declaration; }
	}

	public final Function_declarationContext function_declaration() throws RecognitionException {
		Function_declarationContext _localctx = new Function_declarationContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_function_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			match(T__13);
			setState(106);
			identifier();
			setState(107);
			match(T__2);
			setState(116);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LETTER) {
				{
				setState(108);
				identifier();
				setState(113);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__11) {
					{
					{
					setState(109);
					match(T__11);
					setState(110);
					identifier();
					}
					}
					setState(115);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(118);
			match(T__3);
			setState(119);
			block();
			setState(120);
			match(T__14);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 7529917962L) != 0)) {
				{
				{
				setState(122);
				statement();
				}
				}
				setState(127);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_statementContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public If_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_statement; }
	}

	public final If_statementContext if_statement() throws RecognitionException {
		If_statementContext _localctx = new If_statementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_if_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			match(T__15);
			setState(129);
			expression(0);
			setState(130);
			match(T__16);
			setState(131);
			block();
			setState(139);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__17) {
				{
				{
				setState(132);
				match(T__17);
				setState(133);
				expression(0);
				setState(134);
				match(T__16);
				setState(135);
				block();
				}
				}
				setState(141);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(144);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__18) {
				{
				setState(142);
				match(T__18);
				setState(143);
				block();
				}
			}

			setState(146);
			match(T__14);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class For_statementContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public For_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_statement; }
	}

	public final For_statementContext for_statement() throws RecognitionException {
		For_statementContext _localctx = new For_statementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_for_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			match(T__19);
			setState(149);
			identifier();
			setState(150);
			match(T__1);
			setState(151);
			expression(0);
			setState(152);
			match(T__11);
			setState(153);
			expression(0);
			setState(156);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(154);
				match(T__11);
				setState(155);
				expression(0);
				}
			}

			setState(158);
			match(T__20);
			setState(159);
			block();
			setState(160);
			match(T__14);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class While_statementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public While_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_statement; }
	}

	public final While_statementContext while_statement() throws RecognitionException {
		While_statementContext _localctx = new While_statementContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_while_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			match(T__21);
			setState(163);
			expression(0);
			setState(164);
			match(T__20);
			setState(165);
			block();
			setState(166);
			match(T__14);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TableContext extends ParserRuleContext {
		public List<FieldContext> field() {
			return getRuleContexts(FieldContext.class);
		}
		public FieldContext field(int i) {
			return getRuleContext(FieldContext.class,i);
		}
		public TableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_table; }
	}

	public final TableContext table() throws RecognitionException {
		TableContext _localctx = new TableContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_table);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(168);
			match(T__22);
			setState(177);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 7524593160L) != 0)) {
				{
				setState(169);
				field();
				setState(174);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__11) {
					{
					{
					setState(170);
					match(T__11);
					setState(171);
					field();
					}
					}
					setState(176);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(179);
			match(T__23);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FieldContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public FieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_field; }
	}

	public final FieldContext field() throws RecognitionException {
		FieldContext _localctx = new FieldContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_field);
		try {
			setState(186);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(181);
				identifier();
				setState(182);
				match(T__1);
				setState(183);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(185);
				expression(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Table_accessContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Table_accessContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_table_access; }
	}

	public final Table_accessContext table_access() throws RecognitionException {
		Table_accessContext _localctx = new Table_accessContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_table_access);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(188);
			identifier();
			setState(189);
			match(T__24);
			setState(190);
			expression(0);
			setState(191);
			match(T__25);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Multi_line_commentContext extends ParserRuleContext {
		public Multi_line_commentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_multi_line_comment; }
	}

	public final Multi_line_commentContext multi_line_comment() throws RecognitionException {
		Multi_line_commentContext _localctx = new Multi_line_commentContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_multi_line_comment);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(193);
			match(T__26);
			setState(197);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(194);
					_la = _input.LA(1);
					if ( _la <= 0 || (_la==T__27) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					} 
				}
				setState(199);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
			}
			setState(200);
			match(T__28);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierContext extends ParserRuleContext {
		public List<TerminalNode> LETTER() { return getTokens(lua_grammar_antlr4Parser.LETTER); }
		public TerminalNode LETTER(int i) {
			return getToken(lua_grammar_antlr4Parser.LETTER, i);
		}
		public List<TerminalNode> DIGIT() { return getTokens(lua_grammar_antlr4Parser.DIGIT); }
		public TerminalNode DIGIT(int i) {
			return getToken(lua_grammar_antlr4Parser.DIGIT, i);
		}
		public IdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier; }
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_identifier);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(202);
			match(LETTER);
			setState(206);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(203);
					_la = _input.LA(1);
					if ( !(_la==LETTER || _la==DIGIT) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					} 
				}
				setState(208);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 3:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		case 1:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001\"\u00d2\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0001\u0000\u0001\u0000\u0005\u0000%\b\u0000"+
		"\n\u0000\f\u0000(\t\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0003\u00010\b\u0001\u0001\u0002\u0003\u0002"+
		"3\b\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003C\b\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0005\u0003"+
		"K\b\u0003\n\u0003\f\u0003N\t\u0003\u0001\u0004\u0001\u0004\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005W\b\u0005"+
		"\n\u0005\f\u0005Z\t\u0005\u0003\u0005\\\b\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0003\u0005a\b\u0005\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007p\b\u0007"+
		"\n\u0007\f\u0007s\t\u0007\u0003\u0007u\b\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\b\u0005\b|\b\b\n\b\f\b\u007f\t\b\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0005"+
		"\t\u008a\b\t\n\t\f\t\u008d\t\t\u0001\t\u0001\t\u0003\t\u0091\b\t\u0001"+
		"\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0003\n\u009d\b\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f"+
		"\u0001\f\u0001\f\u0005\f\u00ad\b\f\n\f\f\f\u00b0\t\f\u0003\f\u00b2\b\f"+
		"\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003\r\u00bb"+
		"\b\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000f\u0001\u000f\u0005\u000f\u00c4\b\u000f\n\u000f\f\u000f\u00c7\t\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0005\u0010\u00cd\b\u0010"+
		"\n\u0010\f\u0010\u00d0\t\u0010\u0001\u0010\u0000\u0001\u0006\u0011\u0000"+
		"\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c"+
		"\u001e \u0000\u0005\u0001\u0000\u0005\u0006\u0001\u0000\u0007\b\u0002"+
		"\u0000\t\u000b\u001e\u001f\u0001\u0000\u001c\u001c\u0001\u0000 !\u00dd"+
		"\u0000&\u0001\u0000\u0000\u0000\u0002/\u0001\u0000\u0000\u0000\u00042"+
		"\u0001\u0000\u0000\u0000\u0006B\u0001\u0000\u0000\u0000\bO\u0001\u0000"+
		"\u0000\u0000\n`\u0001\u0000\u0000\u0000\fb\u0001\u0000\u0000\u0000\u000e"+
		"i\u0001\u0000\u0000\u0000\u0010}\u0001\u0000\u0000\u0000\u0012\u0080\u0001"+
		"\u0000\u0000\u0000\u0014\u0094\u0001\u0000\u0000\u0000\u0016\u00a2\u0001"+
		"\u0000\u0000\u0000\u0018\u00a8\u0001\u0000\u0000\u0000\u001a\u00ba\u0001"+
		"\u0000\u0000\u0000\u001c\u00bc\u0001\u0000\u0000\u0000\u001e\u00c1\u0001"+
		"\u0000\u0000\u0000 \u00ca\u0001\u0000\u0000\u0000\"%\u0003\u0002\u0001"+
		"\u0000#%\u0003\u001e\u000f\u0000$\"\u0001\u0000\u0000\u0000$#\u0001\u0000"+
		"\u0000\u0000%(\u0001\u0000\u0000\u0000&$\u0001\u0000\u0000\u0000&\'\u0001"+
		"\u0000\u0000\u0000\'\u0001\u0001\u0000\u0000\u0000(&\u0001\u0000\u0000"+
		"\u0000)0\u0003\u0004\u0002\u0000*0\u0003\u0006\u0003\u0000+0\u0003\u0012"+
		"\t\u0000,0\u0003\u0014\n\u0000-0\u0003\u0016\u000b\u0000.0\u0003\u000e"+
		"\u0007\u0000/)\u0001\u0000\u0000\u0000/*\u0001\u0000\u0000\u0000/+\u0001"+
		"\u0000\u0000\u0000/,\u0001\u0000\u0000\u0000/-\u0001\u0000\u0000\u0000"+
		"/.\u0001\u0000\u0000\u00000\u0003\u0001\u0000\u0000\u000013\u0005\u0001"+
		"\u0000\u000021\u0001\u0000\u0000\u000023\u0001\u0000\u0000\u000034\u0001"+
		"\u0000\u0000\u000045\u0003 \u0010\u000056\u0005\u0002\u0000\u000067\u0003"+
		"\u0006\u0003\u00007\u0005\u0001\u0000\u0000\u000089\u0006\u0003\uffff"+
		"\uffff\u00009C\u0003\b\u0004\u0000:C\u0003 \u0010\u0000;C\u0003\n\u0005"+
		"\u0000<C\u0003\u0018\f\u0000=C\u0003\u001c\u000e\u0000>?\u0005\u0003\u0000"+
		"\u0000?@\u0003\u0006\u0003\u0000@A\u0005\u0004\u0000\u0000AC\u0001\u0000"+
		"\u0000\u0000B8\u0001\u0000\u0000\u0000B:\u0001\u0000\u0000\u0000B;\u0001"+
		"\u0000\u0000\u0000B<\u0001\u0000\u0000\u0000B=\u0001\u0000\u0000\u0000"+
		"B>\u0001\u0000\u0000\u0000CL\u0001\u0000\u0000\u0000DE\n\u0002\u0000\u0000"+
		"EF\u0007\u0000\u0000\u0000FK\u0003\u0006\u0003\u0003GH\n\u0001\u0000\u0000"+
		"HI\u0007\u0001\u0000\u0000IK\u0003\u0006\u0003\u0002JD\u0001\u0000\u0000"+
		"\u0000JG\u0001\u0000\u0000\u0000KN\u0001\u0000\u0000\u0000LJ\u0001\u0000"+
		"\u0000\u0000LM\u0001\u0000\u0000\u0000M\u0007\u0001\u0000\u0000\u0000"+
		"NL\u0001\u0000\u0000\u0000OP\u0007\u0002\u0000\u0000P\t\u0001\u0000\u0000"+
		"\u0000QR\u0003 \u0010\u0000R[\u0005\u0003\u0000\u0000SX\u0003\u0006\u0003"+
		"\u0000TU\u0005\f\u0000\u0000UW\u0003\u0006\u0003\u0000VT\u0001\u0000\u0000"+
		"\u0000WZ\u0001\u0000\u0000\u0000XV\u0001\u0000\u0000\u0000XY\u0001\u0000"+
		"\u0000\u0000Y\\\u0001\u0000\u0000\u0000ZX\u0001\u0000\u0000\u0000[S\u0001"+
		"\u0000\u0000\u0000[\\\u0001\u0000\u0000\u0000\\]\u0001\u0000\u0000\u0000"+
		"]^\u0005\u0004\u0000\u0000^a\u0001\u0000\u0000\u0000_a\u0003\f\u0006\u0000"+
		"`Q\u0001\u0000\u0000\u0000`_\u0001\u0000\u0000\u0000a\u000b\u0001\u0000"+
		"\u0000\u0000bc\u0005\r\u0000\u0000cd\u0005\u0003\u0000\u0000de\u0003 "+
		"\u0010\u0000ef\u0005\f\u0000\u0000fg\u0003\u0006\u0003\u0000gh\u0005\u0004"+
		"\u0000\u0000h\r\u0001\u0000\u0000\u0000ij\u0005\u000e\u0000\u0000jk\u0003"+
		" \u0010\u0000kt\u0005\u0003\u0000\u0000lq\u0003 \u0010\u0000mn\u0005\f"+
		"\u0000\u0000np\u0003 \u0010\u0000om\u0001\u0000\u0000\u0000ps\u0001\u0000"+
		"\u0000\u0000qo\u0001\u0000\u0000\u0000qr\u0001\u0000\u0000\u0000ru\u0001"+
		"\u0000\u0000\u0000sq\u0001\u0000\u0000\u0000tl\u0001\u0000\u0000\u0000"+
		"tu\u0001\u0000\u0000\u0000uv\u0001\u0000\u0000\u0000vw\u0005\u0004\u0000"+
		"\u0000wx\u0003\u0010\b\u0000xy\u0005\u000f\u0000\u0000y\u000f\u0001\u0000"+
		"\u0000\u0000z|\u0003\u0002\u0001\u0000{z\u0001\u0000\u0000\u0000|\u007f"+
		"\u0001\u0000\u0000\u0000}{\u0001\u0000\u0000\u0000}~\u0001\u0000\u0000"+
		"\u0000~\u0011\u0001\u0000\u0000\u0000\u007f}\u0001\u0000\u0000\u0000\u0080"+
		"\u0081\u0005\u0010\u0000\u0000\u0081\u0082\u0003\u0006\u0003\u0000\u0082"+
		"\u0083\u0005\u0011\u0000\u0000\u0083\u008b\u0003\u0010\b\u0000\u0084\u0085"+
		"\u0005\u0012\u0000\u0000\u0085\u0086\u0003\u0006\u0003\u0000\u0086\u0087"+
		"\u0005\u0011\u0000\u0000\u0087\u0088\u0003\u0010\b\u0000\u0088\u008a\u0001"+
		"\u0000\u0000\u0000\u0089\u0084\u0001\u0000\u0000\u0000\u008a\u008d\u0001"+
		"\u0000\u0000\u0000\u008b\u0089\u0001\u0000\u0000\u0000\u008b\u008c\u0001"+
		"\u0000\u0000\u0000\u008c\u0090\u0001\u0000\u0000\u0000\u008d\u008b\u0001"+
		"\u0000\u0000\u0000\u008e\u008f\u0005\u0013\u0000\u0000\u008f\u0091\u0003"+
		"\u0010\b\u0000\u0090\u008e\u0001\u0000\u0000\u0000\u0090\u0091\u0001\u0000"+
		"\u0000\u0000\u0091\u0092\u0001\u0000\u0000\u0000\u0092\u0093\u0005\u000f"+
		"\u0000\u0000\u0093\u0013\u0001\u0000\u0000\u0000\u0094\u0095\u0005\u0014"+
		"\u0000\u0000\u0095\u0096\u0003 \u0010\u0000\u0096\u0097\u0005\u0002\u0000"+
		"\u0000\u0097\u0098\u0003\u0006\u0003\u0000\u0098\u0099\u0005\f\u0000\u0000"+
		"\u0099\u009c\u0003\u0006\u0003\u0000\u009a\u009b\u0005\f\u0000\u0000\u009b"+
		"\u009d\u0003\u0006\u0003\u0000\u009c\u009a\u0001\u0000\u0000\u0000\u009c"+
		"\u009d\u0001\u0000\u0000\u0000\u009d\u009e\u0001\u0000\u0000\u0000\u009e"+
		"\u009f\u0005\u0015\u0000\u0000\u009f\u00a0\u0003\u0010\b\u0000\u00a0\u00a1"+
		"\u0005\u000f\u0000\u0000\u00a1\u0015\u0001\u0000\u0000\u0000\u00a2\u00a3"+
		"\u0005\u0016\u0000\u0000\u00a3\u00a4\u0003\u0006\u0003\u0000\u00a4\u00a5"+
		"\u0005\u0015\u0000\u0000\u00a5\u00a6\u0003\u0010\b\u0000\u00a6\u00a7\u0005"+
		"\u000f\u0000\u0000\u00a7\u0017\u0001\u0000\u0000\u0000\u00a8\u00b1\u0005"+
		"\u0017\u0000\u0000\u00a9\u00ae\u0003\u001a\r\u0000\u00aa\u00ab\u0005\f"+
		"\u0000\u0000\u00ab\u00ad\u0003\u001a\r\u0000\u00ac\u00aa\u0001\u0000\u0000"+
		"\u0000\u00ad\u00b0\u0001\u0000\u0000\u0000\u00ae\u00ac\u0001\u0000\u0000"+
		"\u0000\u00ae\u00af\u0001\u0000\u0000\u0000\u00af\u00b2\u0001\u0000\u0000"+
		"\u0000\u00b0\u00ae\u0001\u0000\u0000\u0000\u00b1\u00a9\u0001\u0000\u0000"+
		"\u0000\u00b1\u00b2\u0001\u0000\u0000\u0000\u00b2\u00b3\u0001\u0000\u0000"+
		"\u0000\u00b3\u00b4\u0005\u0018\u0000\u0000\u00b4\u0019\u0001\u0000\u0000"+
		"\u0000\u00b5\u00b6\u0003 \u0010\u0000\u00b6\u00b7\u0005\u0002\u0000\u0000"+
		"\u00b7\u00b8\u0003\u0006\u0003\u0000\u00b8\u00bb\u0001\u0000\u0000\u0000"+
		"\u00b9\u00bb\u0003\u0006\u0003\u0000\u00ba\u00b5\u0001\u0000\u0000\u0000"+
		"\u00ba\u00b9\u0001\u0000\u0000\u0000\u00bb\u001b\u0001\u0000\u0000\u0000"+
		"\u00bc\u00bd\u0003 \u0010\u0000\u00bd\u00be\u0005\u0019\u0000\u0000\u00be"+
		"\u00bf\u0003\u0006\u0003\u0000\u00bf\u00c0\u0005\u001a\u0000\u0000\u00c0"+
		"\u001d\u0001\u0000\u0000\u0000\u00c1\u00c5\u0005\u001b\u0000\u0000\u00c2"+
		"\u00c4\b\u0003\u0000\u0000\u00c3\u00c2\u0001\u0000\u0000\u0000\u00c4\u00c7"+
		"\u0001\u0000\u0000\u0000\u00c5\u00c3\u0001\u0000\u0000\u0000\u00c5\u00c6"+
		"\u0001\u0000\u0000\u0000\u00c6\u00c8\u0001\u0000\u0000\u0000\u00c7\u00c5"+
		"\u0001\u0000\u0000\u0000\u00c8\u00c9\u0005\u001d\u0000\u0000\u00c9\u001f"+
		"\u0001\u0000\u0000\u0000\u00ca\u00ce\u0005 \u0000\u0000\u00cb\u00cd\u0007"+
		"\u0004\u0000\u0000\u00cc\u00cb\u0001\u0000\u0000\u0000\u00cd\u00d0\u0001"+
		"\u0000\u0000\u0000\u00ce\u00cc\u0001\u0000\u0000\u0000\u00ce\u00cf\u0001"+
		"\u0000\u0000\u0000\u00cf!\u0001\u0000\u0000\u0000\u00d0\u00ce\u0001\u0000"+
		"\u0000\u0000\u0015$&/2BJLX[`qt}\u008b\u0090\u009c\u00ae\u00b1\u00ba\u00c5"+
		"\u00ce";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}